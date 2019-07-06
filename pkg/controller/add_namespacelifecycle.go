package controller

import (
	"github.com/owenn/project-lifecycle/pkg/controller/namespacelifecycle"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, namespacelifecycle.Add)
}
