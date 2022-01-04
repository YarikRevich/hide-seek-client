package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/YarikRevich/hide-seek-client/assets"
	"github.com/YarikRevich/hide-seek-client/internal/core/middlewares"
	"github.com/YarikRevich/hide-seek-client/internal/core/paths"
	"github.com/YarikRevich/hide-seek-client/internal/core/profiling"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/loop"

	"github.com/YarikRevich/hide-seek-client/tools/params"
	"github.com/YarikRevich/hide-seek-client/tools/printer"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	rand.Seed(time.Now().Unix())

	flag.Parse()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	lgf, err := os.OpenFile(filepath.Join(paths.GAME_LOG_DIR, "/log.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	logrus.SetOutput(lgf)

	if params.IsDebug() {
		logrus.SetLevel(logrus.DebugLevel)
		profiling.UseProfiler().Init()
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	sources.UseSources().LoadSources(assets.Assets)

	middlewares.UseMiddlewares().Prepare().Use()

	printer.PrintCliMessage("HideSeek\nClient!")
	if params.IsDebug() {
		printer.PrintCliMessage("You are in\nDebug mode")
	}

	debug.SetGCPercent(2000)
}

func main() {

	s := screen.UseScreen()
	ebiten.SetWindowSize(s.GetMaxWidth(), s.GetMaxHeight())
	ebiten.SetWindowTitle("HideSeek-Client")
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSizeLimits(s.GetMinWidth(), s.GetMinHeight(), -1, -1)

	log.Fatalln(ebiten.RunGame(loop.New()))
}