package analyze

import (
	"github.com/fossas/fossa-cli/analyze/module"
)

// We should get rid of the indirection of module.Type and just have this return (module.Analyzer, error) where error can be ErrNotFound
func New(t module.Type) module.Analyzer {}
