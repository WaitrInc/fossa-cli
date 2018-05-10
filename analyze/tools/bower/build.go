package bower

import "github.com/fossas/fossa-cli/analyze/module"

// Build makes a best-effort attempt at building the module.
func (a *Analyzer) Build(m module.Module) error {}

// IsBuilt checks whether a module has been built.
func (a *Analyzer) IsBuilt(m module.Module) (bool, error) {}
