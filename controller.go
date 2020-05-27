package main

import (
	clientset "github.com/faizan82/controllers/cnat-controller/pkg/generated/clientset/versioned"
)

// Controller is our reconciler that brings system to the desired state
type Controller struct {
	cnatClientSet clientset.Interface
}