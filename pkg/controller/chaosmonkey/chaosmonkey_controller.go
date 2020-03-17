package chaosmonkey

import (
	"context"
	"time"

	iov1alpha1 "github.com/lauripiispanen/chaos-operator/pkg/apis/io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_chaosmonkey")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ChaosMonkey Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileChaosMonkey{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	controllerName := "chaosmonkey-controller"
	c, err := controller.New(controllerName, mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ChaosMonkey
	err = c.Watch(&source.Kind{Type: &iov1alpha1.ChaosMonkey{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	mapFn := handler.ToRequestsFunc(func(a handler.MapObject) []reconcile.Request {
		return []reconcile.Request{
			{NamespacedName: types.NamespacedName{
				Name:      controllerName,
				Namespace: a.Meta.GetNamespace(),
			}},
		}
	})

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ChaosMonkey
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestsFromMapFunc{
		ToRequests: mapFn,
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileChaosMonkey implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileChaosMonkey{}

// ReconcileChaosMonkey reconciles a ChaosMonkey object
type ReconcileChaosMonkey struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ChaosMonkey object and makes changes based on the state read
// and what is in the ChaosMonkey.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileChaosMonkey) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ChaosMonkey")

	// Fetch the ChaosMonkey instance
	instance := &iov1alpha1.ChaosMonkey{}
	ctx := context.TODO()
	err := r.client.Get(ctx, request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// ChaosMonkey object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// ChaosMonkey object was found, check if we've exceeded last run time + interval and if we have, delete a pod
	if instance.Status.LastRunTime < time.Now().Unix()-instance.Spec.Interval {
		podList := &corev1.PodList{}
		opts := []client.ListOption{
			client.InNamespace(request.NamespacedName.Namespace),
			client.MatchingFields{"status.phase": "Running"},
		}
		err := r.client.List(ctx, podList, opts...)
		if err != nil {
			return reconcile.Result{}, err
		}
		var deletablePod *corev1.Pod
		for _, pod := range podList.Items {
			if pod.Name != request.Name {
				deletablePod = &pod
				break
			}
		}
		if deletablePod != nil {
			reqLogger.Info("Deleting pod...", "Pod.Namespace", deletablePod.Namespace, "Pod.Name", deletablePod.Name)
			if err := r.client.Delete(ctx, deletablePod); err != nil {
				return reconcile.Result{}, err
			}
			instance.Status.LastRunTime = time.Now().Unix()
			if err := r.client.Status().Update(ctx, instance); err != nil {
				return reconcile.Result{}, err
			}
		}
	}

	return reconcile.Result{RequeueAfter: time.Second * time.Duration(instance.Spec.Interval)}, nil
}
