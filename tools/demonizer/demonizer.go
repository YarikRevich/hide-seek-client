package demonizer

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
)

type Demonizer struct {
	lockWord string
}

func (d *Demonizer) Demonize() {
	if _, ok := os.LookupEnv(d.lockWord); !ok {
		var args []string
		if len(os.Args) > 1 {
			args = os.Args[0:]
		} else {
			args = os.Args
		}
		p, err := os.Executable()
		if err != nil {
			logrus.Fatal(err)
		}
		command := exec.Command(p, args...)
		command.Env = append(command.Env, os.Environ()...)
		command.Env = append(command.Env, d.lockWord)

		_, err = syscall.Setsid()
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func New(lockWord string) *Demonizer {
	return &Demonizer{lockWord}
}
