/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	testappsv1 "test-apps/pending-handler/api/v1"
)

// PendingHandlerReconciler reconciles a PendingHandler object
type PendingHandlerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=test-apps.test-apps,resources=pendinghandlers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=test-apps.test-apps,resources=pendinghandlers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=test-apps.test-apps,resources=pendinghandlers/finalizers,verbs=update

// +kubebuilder:rbac:groups="",resources=pods,verbs=get;delete
// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PendingHandler object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *PendingHandlerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Algorithm for this controller
	// 1) Get PendingSpec
	// 2) Get a list of pods (limit to specific namespaces for now)
	// 3) If Pod has matching spec than we will delete the pod.

	var pendingAPI testappsv1.PendingHandler
	if err := r.Get(ctx, req.NamespacedName, &pendingAPI); err != nil {
		log.Error(err, "unable to fetch PendingHandler")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	// Retrieve the pods in the specified namespace
	var pods core.PodList
	if err := r.List(ctx, &pods, client.InNamespace(pendingAPI.Spec.Namespace)); err != nil {
		log.Error(err, "unable to list pods")
		return ctrl.Result{}, err
	}
	for _, val := range pods.Items {
		log.Info("pod name list", "podName", val.Name)
		podCopy := val.DeepCopy()
		containerStatus := podCopy.Status.ContainerStatuses
		podConditions := podCopy.Status.Conditions

		if matchContainerStatusFromPendingHandler(containerStatus, pendingAPI.Spec.PendingContainerStatuses) {
			log.Info("pod pending handler detected stuck pod in condition", "pod to be deleted", val.Name)
			if err := r.Delete(ctx, podCopy); err != nil {
				log.Error(err, "unable to delete pod", podCopy.Name)
			}
		}
		if matchConditionFromPendingHandler(podConditions, pendingAPI.Spec.PendingConditions) {
			log.Info("pod pending handler detected stuck pod in container status", "pod to be deleted", val.Name)
			if err := r.Delete(ctx, podCopy); err != nil {
				log.Error(err, "unable to delete pod", podCopy.Name)
			}
		}

	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PendingHandlerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testappsv1.PendingHandler{}).
		Complete(r)
}

func matchContainerStatusFromPendingHandler(cStatus []core.ContainerStatus, pendingContainterStatuses []testappsv1.PendingContainerStatusEntry) bool {
	for _, val := range cStatus {
		for _, match := range pendingContainterStatuses {
			if val.State.Waiting.Reason == match.Condition {
				return true
			}
		}
	}
	return false
}

func matchConditionFromPendingHandler(cStatus []core.PodCondition, pendingConditions []testappsv1.PendingConditionsEntry) bool {
	for _, val := range cStatus {
		for _, match := range pendingConditions {
			if string(val.Type) == match.Condition {
				return true
			}
		}
	}
	return false
}
