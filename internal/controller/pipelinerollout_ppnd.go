package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	numaflowv1 "github.com/numaproj/numaflow/pkg/apis/numaflow/v1alpha1"
	"github.com/numaproj/numaplane/internal/common"
	"github.com/numaproj/numaplane/internal/util/kubernetes"
	"github.com/numaproj/numaplane/internal/util/logger"
	apiv1 "github.com/numaproj/numaplane/pkg/apis/numaplane/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8stypes "k8s.io/apimachinery/pkg/types"
)

const (
	// unpausibleDelay represents the amount of time that we will allow before a Pipeline is considered "unpausible", determined if the pipeline is in a "Failed" state and has had
	// a desiredPhase=Paused for this long
	unpausibleDelay time.Duration = 30 * time.Second
)

// TODO: move PPND logic out to its own separate file
// normal sequence of events when we need to pause:
// - set Pipeline's desiredPhase=Paused
// - wait for the desire to Pause to be reconciled completely
// - if we need to update the Pipeline spec:
//   - update it
//   - wait for the spec update to be reconciled completely
//
// - as long as there's no other requirement to pause, set desiredPhase=Running
// return boolean for whether we can stop the PPND process
func (r *PipelineRolloutReconciler) processExistingPipelineWithPPND(ctx context.Context, pipelineRollout *apiv1.PipelineRollout,
	existingPipelineDef, newPipelineDef *kubernetes.GenericObject) (bool, error) {

	numaLogger := logger.FromContext(ctx)

	var newPipelineSpec PipelineSpec
	if err := json.Unmarshal(newPipelineDef.Spec.Raw, &newPipelineSpec); err != nil {
		return false, fmt.Errorf("failed to convert new Pipeline spec %q into PipelineSpec type, err=%v", string(newPipelineDef.Spec.Raw), err)
	}

	pipelineNeedsToUpdate, err := pipelineSpecNeedsUpdating(ctx, existingPipelineDef, newPipelineDef)
	if err != nil {
		return false, err
	}

	needsPaused, err := r.shouldBePaused(ctx, pipelineRollout, existingPipelineDef, newPipelineDef, pipelineNeedsToUpdate)
	if err != nil {
		return false, err
	}
	if needsPaused == nil { // not enough information to know
		return false, errors.New("not enough information available to know if we need to pause")
	}
	shouldBePaused := *needsPaused

	pipelineUnpausible := r.checkPipelineUnpausible(ctx, shouldBePaused, existingPipelineDef, pipelineRollout)
	shouldBePaused = shouldBePaused && !pipelineUnpausible

	if err := r.setPipelineLifecycle(ctx, shouldBePaused, existingPipelineDef); err != nil {
		return false, err
	}

	pipelinePaused := checkPipelineStatus(ctx, existingPipelineDef, numaflowv1.PipelinePhasePaused)

	// update the ResourceVersion in the newPipelineDef in case it got updated
	newPipelineDef.ResourceVersion = existingPipelineDef.ResourceVersion

	// if it's safe to Update and we need to, do it now
	if pipelineNeedsToUpdate {
		pipelineRollout.Status.MarkPending()

		if !shouldBePaused || (shouldBePaused && pipelinePaused) {
			numaLogger.Infof("it's safe to update Pipeline so updating now")
			r.recorder.Eventf(pipelineRollout, "Normal", "PipelineUpdate", "it's safe to update Pipeline so updating now")

			if shouldBePaused {
				err = withDesiredPhase(newPipelineDef, "Paused")
				if err != nil {
					return false, err
				}
			}
			err = kubernetes.UpdateCR(ctx, r.restConfig, newPipelineDef, "pipelines")
			if err != nil {
				return false, err
			}
			pipelineRollout.Status.MarkDeployed(pipelineRollout.Generation)
		}
	} else {
		pipelineRollout.Status.MarkDeployed(pipelineRollout.Generation)
	}

	// are we done with PPND?
	doneWithPPND := !shouldBePaused

	// but if the PipelineRollout says to pause and we're Paused, this is also "doneWithPPND"
	specBasedPause := newPipelineSpec.Lifecycle.DesiredPhase == string(numaflowv1.PipelinePhasePaused) || newPipelineSpec.Lifecycle.DesiredPhase == string(numaflowv1.PipelinePhasePausing)
	if specBasedPause && (pipelinePaused || pipelineUnpausible) {
		doneWithPPND = true
	}

	return doneWithPPND, nil
}

// Does the Pipeline need to be paused?
// This is based on:
//
//	any difference in spec between PipelineRollout and Pipeline, with the exception of lifecycle.desiredPhase field
//	any pause request coming from isbsvc or Numaflow Controller
//	spec says to pause
//	pipeline spec change is still being reconciled
//
// return whether to pause, not to pause, or otherwise unknown
func (r *PipelineRolloutReconciler) shouldBePaused(ctx context.Context, pipelineRollout *apiv1.PipelineRollout, existingPipelineDef, newPipelineDef *kubernetes.GenericObject, pipelineNeedsToUpdate bool) (*bool, error) {
	numaLogger := logger.FromContext(ctx)

	var newPipelineSpec PipelineSpec
	if err := json.Unmarshal(newPipelineDef.Spec.Raw, &newPipelineSpec); err != nil {
		return nil, fmt.Errorf("failed to convert new Pipeline spec %q into PipelineSpec type, err=%v", string(newPipelineDef.Spec.Raw), err)
	}
	var existingPipelineSpec PipelineSpec
	if err := json.Unmarshal(existingPipelineDef.Spec.Raw, &existingPipelineSpec); err != nil {
		return nil, fmt.Errorf("failed to convert existing Pipeline spec %q into PipelineSpec type, err=%v", string(existingPipelineDef.Spec.Raw), err)
	}

	// is the Pipeline currently being reconciled?
	pipelineUpdating, err := pipelineIsUpdating(newPipelineDef, existingPipelineDef)
	if err != nil {
		return nil, err
	}

	// is the Pipeline currently being reconciled while our desiredPhase==Paused?
	// only in this circumstance do we need to make sure to remain Paused until that reconciliation is complete
	existingPipelinePauseDesired := existingPipelineSpec.Lifecycle.DesiredPhase == string(numaflowv1.PipelinePhasePaused)
	pipelineUpdating = pipelineUpdating && existingPipelinePauseDesired

	// Is either Numaflow Controller or ISBService trying to update (such that we need to pause)?
	externalPauseRequest, pauseRequestsKnown, err := r.checkForPauseRequest(ctx, pipelineRollout, getISBSvcName(newPipelineSpec))
	if err != nil {
		return nil, err
	}

	// check to see if the PipelineRollout spec itself says to Pause
	specBasedPause := (newPipelineSpec.Lifecycle.DesiredPhase == string(numaflowv1.PipelinePhasePaused) || newPipelineSpec.Lifecycle.DesiredPhase == string(numaflowv1.PipelinePhasePausing))

	shouldBePaused := pipelineNeedsToUpdate || pipelineUpdating || externalPauseRequest || specBasedPause
	numaLogger.Debugf("shouldBePaused=%t, pipelineNeedsToUpdate=%t, pipelineUpdating=%t, externalPauseRequest=%t, specBasedPause=%t",
		shouldBePaused, pipelineNeedsToUpdate, pipelineUpdating, externalPauseRequest, specBasedPause)

	// if we have incomplete pause request information (i.e. numaflowcontrollerrollout or isbservicerollout not yet reconciled), don't return
	// that it's okay to run
	if !shouldBePaused && !pauseRequestsKnown {
		numaLogger.Debugf("incomplete pause request information")
		return nil, nil
	}

	return &shouldBePaused, nil
}

// determine if Pipeline has been trying to pause for a period of time and is in "Failed" state
// determine what the Pipeline annotation should be, and set it if it's not already set
func (r *PipelineRolloutReconciler) checkPipelineUnpausible(ctx context.Context, shouldBePaused bool, existingPipelineDef *kubernetes.GenericObject, pipelineRollout *apiv1.PipelineRollout) bool {
	if shouldBePaused {
		// do we need to set DesiredPauseBeginTime?
		if pipelineRollout.Status.PPNDStatus.DesiredPauseBeginTime.IsZero() {
			pipelineRollout.Status.PPNDStatus.DesiredPauseBeginTime = metav1.NewTime(time.Now())
			r.enqueuePipelineRolloutWithDelay(k8stypes.NamespacedName{Namespace: pipelineRollout.Namespace, Name: pipelineRollout.Name}, unpausibleDelay+1*time.Second)
			return false
		}

		// if Failed and 30s since desiredPausedStartTime
		if checkPipelineStatus(ctx, existingPipelineDef, numaflowv1.PipelinePhaseFailed) && time.Since(pipelineRollout.Status.PPNDStatus.DesiredPauseBeginTime.Time) > unpausibleDelay {
			// TODO: set unpausible annotation
			return true
		}
		return false
	} else {
		// clear DesiredPauseBeginTime
		pipelineRollout.Status.PPNDStatus.DesiredPauseBeginTime = metav1.NewTime(time.Time{})
		// TODO: clear annotation on pipeline
		return false
	}
}

// do we need to start the PPND process, if we haven't already?
// this is based on if:
//
//	there's any difference in spec between PipelineRollout and Pipeline
//	any pause request coming from isbsvc or Numaflow Controller
func (r *PipelineRolloutReconciler) needPPND(ctx context.Context, pipelineRollout *apiv1.PipelineRollout, newPipelineDef *kubernetes.GenericObject, pipelineUpdateRequiringPPND bool) (*bool, error) {
	numaLogger := logger.FromContext(ctx)

	var newPipelineSpec PipelineSpec
	if err := json.Unmarshal(newPipelineDef.Spec.Raw, &newPipelineSpec); err != nil {
		return nil, fmt.Errorf("failed to convert new Pipeline spec %q into PipelineSpec type, err=%v", string(newPipelineDef.Spec.Raw), err)
	}

	// Is either Numaflow Controller or ISBService trying to update (such that we need to pause)?
	externalPauseRequest, pauseRequestsKnown, err := r.checkForPauseRequest(ctx, pipelineRollout, getISBSvcName(newPipelineSpec))
	if err != nil {
		return nil, err
	}

	needPPND := externalPauseRequest || pipelineUpdateRequiringPPND
	numaLogger.Debugf("needPPND=%t, externalPauseRequest=%t, pipelineUpdateRequiringPPND=%t", needPPND, externalPauseRequest, pipelineUpdateRequiringPPND)

	// if we have incomplete pause request information (i.e. numaflowcontrollerrollout or isbservicerollout not yet reconciled), don't return
	// that it's okay not to pause because we don't know for sure
	if !needPPND && !pauseRequestsKnown {
		numaLogger.Debugf("incomplete pause request information")
		return nil, nil
	}

	return &needPPND, nil
}

// check for all pause requests for this Pipeline (i.e. both from Numaflow Controller and ISBService)
// return:
// - whether there's a pause request
// - whether all pause requests are known
// - error if any
func (r *PipelineRolloutReconciler) checkForPauseRequest(ctx context.Context, pipelineRollout *apiv1.PipelineRollout, isbsvcName string) (bool, bool, error) {
	numaLogger := logger.FromContext(ctx)

	pm := GetPauseModule()

	// Is either Numaflow Controller or ISBService trying to update (such that we need to pause)?
	controllerPauseRequest, found := pm.getPauseRequest(pm.getNumaflowControllerKey(pipelineRollout.Namespace))
	if !found {
		numaLogger.Debugf("No pause request found for numaflow controller on namespace %q", pipelineRollout.Namespace)
		return false, false, nil

	}
	controllerRequestsPause := controllerPauseRequest != nil && *controllerPauseRequest

	isbsvcPauseRequest, found := pm.getPauseRequest(pm.getISBServiceKey(pipelineRollout.Namespace, isbsvcName))
	if !found {
		numaLogger.Debugf("No pause request found for isbsvc %q on namespace %q", isbsvcName, pipelineRollout.Namespace)
		return false, false, nil
	}
	isbsvcRequestsPause := (isbsvcPauseRequest != nil && *isbsvcPauseRequest)

	return controllerRequestsPause || isbsvcRequestsPause, true, nil
}

// make sure our Pipeline's Lifecycle is what we need it to be
func (r *PipelineRolloutReconciler) setPipelineLifecycle(ctx context.Context, pause bool, existingPipelineDef *kubernetes.GenericObject) error {
	numaLogger := logger.FromContext(ctx)
	var existingPipelineSpec PipelineSpec
	if err := json.Unmarshal(existingPipelineDef.Spec.Raw, &existingPipelineSpec); err != nil {
		return err
	}
	lifeCycleIsPaused := existingPipelineSpec.Lifecycle.DesiredPhase == string(numaflowv1.PipelinePhasePaused)

	if pause && !lifeCycleIsPaused {
		numaLogger.Info("pausing pipeline")
		r.recorder.Eventf(existingPipelineDef, "Normal", "PipelinePause", "pausing pipeline")
		if err := GetPauseModule().pausePipeline(ctx, r.restConfig, existingPipelineDef); err != nil {
			return err
		}
	} else if !pause && lifeCycleIsPaused {
		numaLogger.Info("resuming pipeline")
		r.recorder.Eventf(existingPipelineDef, "Normal", "PipelineResume", "resuming pipeline")

		run, err := GetPauseModule().runPipelineIfSafe(ctx, r.restConfig, existingPipelineDef)
		if err != nil {
			return err
		}
		if !run {
			numaLogger.Infof("new pause request, can't resume pipeline at this time, will try again later")
			r.recorder.Eventf(existingPipelineDef, "Normal", "PipelineResumeFailed", "new pause request, can't resume pipeline at this time, will try again later")
		}
	}
	return nil
}

// return true if Pipeline (or its children) is still in the process of being reconciled
func pipelineIsUpdating(newPipelineDef *kubernetes.GenericObject, existingPipelineDef *kubernetes.GenericObject) (bool, error) {
	existingPipelineStatus, err := kubernetes.ParseStatus(existingPipelineDef)
	if err != nil {
		return false, err
	}
	// if Pipeline's ObservedGeneration is old, then Numaflow Controller hasn't even seen the generation change yet
	if !pipelineObservedGenerationCurrent(newPipelineDef.Generation, existingPipelineStatus.ObservedGeneration) {
		return true, nil
	}

	// note if Pipeline's children are still being updated
	unhealthyOrProgressing, _ := checkChildResources(existingPipelineStatus.Conditions, func(c metav1.Condition) bool {
		return c.Status == metav1.ConditionFalse
	})

	return unhealthyOrProgressing, nil

}

// TODO: consider patch?
func (r *PipelineRolloutReconciler) markPipelineUnpausible(ctx context.Context, mark bool, pipeline *kubernetes.GenericObject) error {
	if mark {
		pipeline.Annotations[common.LabelKeyPipelineUnpausible] = "true"
	} else {
		delete(pipeline.Annotations, common.LabelKeyPipelineUnpausible)
	}
	return kubernetes.UpdateCR(ctx, r.restConfig, pipeline, "pipelines")
}
