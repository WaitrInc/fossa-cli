package bower

import "github.com/fossas/fossa-cli/services"

type Analyzer struct {
	io services.Services
}

func New(io services.Services) (Analyzer, error) {}
