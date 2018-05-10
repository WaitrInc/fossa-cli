package module

import (
	"errors"
)

// ErrNotImplemented is returned whenever an analyzer does not implement a method.
var ErrNotImplemented = errors.New("not implemented")

// An Analyzer is an implementation of functionality for different build systems.
type Analyzer interface {
	// Discover autoconfigures modules in a given directory.
	Discover(dir string) ([]Module, error)

	// Build makes a best-effort attempt at building the module.
	Build(m Module) error
	// IsBuilt checks whether a module has been built.
	IsBuilt(m Module) (bool, error)

	// Analyze runs an analysis of a module.
	Analyze(m Module) (Analyzed, error)
}

// An Analyzed module has dependencies and metadata.
type Analyzed struct {
	Module       Module
	Imports      []PackageID  // Imports are direct dependencies.
	Dependencies []Dependency // Dependencies includes all transitive dependencies.
	Metadata     interface{}
}

// A Dependency is a package. It may itself have dependencies.
type Dependency struct {
	PackageID
	Imports []PackageID
}
