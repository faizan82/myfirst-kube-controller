package main

import (
	clientset "github.com/faizan82/controllers/cnat-controller/pkg/generated/clientset/versioned"
	listers "github.com/faizan82/controllers/cnat-controller/pkg/generated/listers/cnat/v1alpha1"
	"k8s.io/client-go/1.5/kubernetes"
	corev1lister "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

// Controller is our reconciler that brings system to the desired state
type Controller struct {
	cnatClientset clientset.Interface
	kubeClientset kubernetes.Interface

	atListers listers.AtLister
	atsSynced cache.InformerSynced

	podListers corev1lister.PodLister

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}
