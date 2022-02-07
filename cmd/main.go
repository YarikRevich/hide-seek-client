package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/paths"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling/runtime"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/loop"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"

	"github.com/YarikRevich/hide-seek-client/tools/params"
	"github.com/YarikRevich/hide-seek-client/tools/printer"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	sr := beep.SampleRate(44100)
	if err := speaker.Init(sr, sr.N(time.Second/10000)); err != nil {
		logrus.Fatal("error happened initializing audio speaker")
	}

	rand.Seed(time.Now().Unix())

	flag.Parse()

	paths.InitSystemPaths()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	lgf, err := os.OpenFile(filepath.Join(paths.GAME_LOG_DIR, "log.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	logrus.SetOutput(lgf)

	logrus.SetLevel(logrus.WarnLevel)
	if params.IsDebug() {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// sources.UseSources().LoadSources(assets.Assets)

	// middlewares.UseMiddlewares().Prepare().Use()

	printer.PrintCliMessage("HideSeek\nClient!")
	if params.IsDebug() {
		printer.PrintCliMessage("You are in\nDebug mode")
	}

	// debug.SetGCPercent(2000)

}

func main() {
	if params.IsProfileCPU() {
		runtime.UseProfiler().StartMonitoring()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				runtime.UseProfiler().StopMonitoring()
				os.Exit(0)
			}
		}()
	}

	sm := new(screen.ScreenManager)
	maxSize := sm.GetMaxSize()
	minSize := sm.GetMinSize()

	ebiten.SetWindowSize(int(maxSize.X), int(maxSize.Y))
	ebiten.SetWindowTitle("HideSeek-Client")
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSizeLimits(int(minSize.X), int(minSize.Y), -1, -1)

	log.Fatalln(ebiten.RunGame(loop.New(sm)))
}
