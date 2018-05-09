package exec

import "github.com/fossas/fossa-cli/log"

type ExecOptions struct {
	Dir string
	Env map[string]string // defaults to os.Env()
}

func Exec(cmd string, options ExecOptions) error {
	log.Debug()
	// do logging and stuff
}
