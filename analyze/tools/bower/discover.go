package bower

import "github.com/fossas/fossa-cli/analyze/module"

// Discover autoconfigures modules in a given directory.
func (a *Analyzer) Discover(dir string) ([]module.Module, error) {}
