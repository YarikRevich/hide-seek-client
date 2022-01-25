package runtime

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/YarikRevich/hide-seek-client/internal/core/paths"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/common"
	"github.com/sirupsen/logrus"
)

var instance common.Profiler

type Profiler struct {
	file *os.File
}

func (p *Profiler) StartMonitoring(profiler ...string) {
	if err := pprof.StartCPUProfile(p.file); err != nil {
		logrus.Fatalln("could not start CPU profile: ", err)
	}
}
func (p *Profiler) StopMonitoring(profiler ...string) {
	pprof.StopCPUProfile()
}
func (p *Profiler) Show() string {
	return ""
}

func UseProfiler() common.Profiler {
	if instance == nil {
		f, err := os.Create(fmt.Sprintf("%s/%s", paths.GAME_PPROF_DIR, "profilecpu.out"))
		if err != nil {
			logrus.Fatalln("could not create CPU profile: ", err)
		}
		instance = &Profiler{
			file: f,
		}

	}
	return instance
}
