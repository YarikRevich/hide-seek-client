package networking

import (
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
)

type LANServer struct {
	pid int
}

func (l *LANServer) Start() {
	cmd := exec.Command("HideSeek-Server")
	if err := cmd.Start(); err != nil {
		logrus.Fatal(err)
	}
	l.pid = cmd.Process.Pid
}

func (l *LANServer) Stop() {
	if l.pid != 0 {
		if err := syscall.Kill(l.pid, syscall.SIGINT); err != nil {
			logrus.Fatal(err)
		}
	}
}

func NewLANServer() *LANServer {
	return new(LANServer)
}
