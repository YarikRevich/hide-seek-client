package lanserver

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

var instance *LANServer

type LANServer struct {
	process *os.Process
}

func (l *LANServer) Start() {
	cmd := exec.Command("HideSeek-Server")
	if err := cmd.Start(); err != nil {
		logrus.Fatal(err)
	}
	l.process = cmd.Process
}

func (l *LANServer) Stop() {
	fmt.Println(l.process.Kill())
}

func UseLANServer() *LANServer {
	if instance == nil {
		instance = new(LANServer)
	}
	return instance
}
