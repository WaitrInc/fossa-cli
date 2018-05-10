package module

import (
	"fmt"
)

// A Module is a single unit of code (e.g. an entrypoint) to analyze.
type Module struct {
	Name   string `yaml:"name"`
	Target string `yaml:"target"`
	Path   string `yaml:"path"`
	Type   string `yaml:"type"`

	Options Options `yaml:"options"`
}

// Options specifies optional modifiers for module analysis.
type Options struct {
	AllowUnbuilt bool

	Go     GoOptions
	Python PythonOptions
	Gradle GradleOptions
	Maven  MavenOptions
}

// GoOptions specify options for Go modules.
type GoOptions struct {
	AllowUnresolved bool
}

// PythonOptions specify options for Python modules.
type PythonOptions struct {
	Strategy string // pipdeptree or requirements.txt
}

// RubyOptions specify options for Ruby modules.
type RubyOptions struct {
	Strategy string // Gemfile.lock or `bundler list`
}

// We also need strategy options for Bower and Node (output vs. fs parsing)

// GradleOptions specify options for Gradle modules.
type GradleOptions struct {
	Task          string
	Configuration string
}

// MavenOptions specify options for Maven modules.
type MavenOptions struct {
	SettingsFile string
}

// Type is an enumeration of supported build system types
type Type string

const (
	// ----------------
	// Individual tools
	// ----------------

	// Bower is the module type for bower.io
	Bower = Type("bower")
	// Composer is the module type for getcomposer.org
	Composer = Type("composer")
	// Maven is the module type for maven.apache.org
	Maven = Type("maven")
	// SBT is the module type for scala-sbt.org
	SBT = Type("sbt")
	// Gradle is the module type for gradle.org
	Gradle = Type("gradle")
	// NuGet is the module type for nuget.org
	NuGet = Type("nuget")
	// Pip is the module type for https://pip.pypa.io/en/stable/
	Pip = Type("pip")

	// --------------------------------------------
	// Ecosystems where many tools behave similarly
	// --------------------------------------------

	// Ruby is the module type for Bundler (bundler.io)
	Ruby = Type("ruby")
	// Nodejs is the module type for NPM (npmjs.org) and Yarn (yarnpkg.com)
	Nodejs = Type("nodejs")
	// Golang is the module type for dep, glide, godep, govendor, vndr, and manual gopath vendoring
	Golang = Type("golang")
)

// Types holds the list of all available module types for analysis
var Types = []Type{
	Bower,
	Composer,
	Maven,
	SBT,
	Gradle,
	NuGet,
	Pip,
	Ruby,
	Nodejs,
	Golang,
}

// Parse returns a module Type given a string
func Parse(key string) (Type, error) {
	switch key {
	// Bower aliases
	case "bower":
		return Bower, nil

	// Node aliases
	case "commonjspackage":
		fallthrough
	case "npm":
		fallthrough
	case "node":
		fallthrough
	case "nodejs":
		return Nodejs, nil

	// Compower aliases
	case "composer":
		return Composer, nil

	// Golang aliases
	case "gopackage":
		fallthrough
	case "golang":
		fallthrough
	case "go":
		return Golang, nil

	// Maven aliases
	case "maven":
		fallthrough
	case "mvn":
		return Maven, nil

	// Python aliases:
	case "python":
		fallthrough
	case "py":
		fallthrough
	case "pip":
		return Pip, nil

	// Ruby aliases
	case "bundler":
		fallthrough
	case "gem":
		fallthrough
	case "rubygems":
		fallthrough
	case "ruby":
		return Ruby, nil

	// SBT aliases
	case "scala":
		fallthrough
	case "sbtpackage":
		fallthrough
	case "sbt":
		return SBT, nil

	// Gradle aliases
	case "gradle":
		return Gradle, nil

	// NuGet aliases
	case "dotnet":
		fallthrough
	case "csharp":
		fallthrough
	case "nuget":
		return NuGet, nil
	}
	return "", fmt.Errorf("unknown module type: %s", key)
}
