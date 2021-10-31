package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"embed"

	// "github.com/YarikRevich/HideSeek-Client/internal/ai/collisions"
	"github.com/YarikRevich/HideSeek-Client/internal/loop"
	"github.com/YarikRevich/HideSeek-Client/internal/profiling"

	"github.com/YarikRevich/HideSeek-Client/internal/core/paths"
	"github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
	"github.com/YarikRevich/HideSeek-Client/tools/printer"
	"github.com/sirupsen/logrus"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed assets
	assets embed.FS
)

func init() {
	flag.Parse()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	lgf, err := os.OpenFile(filepath.Join(paths.GAME_LOG_DIR, "/log.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	logrus.SetOutput(lgf)

	if cli.GetDebug() {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	resource_manager.LoadResources(map[resource_manager.Component][]resource_manager.Loader{
		{Embed: assets, Path: "assets/images"}: {
			imageloader.Load,
		},
		{Embed: assets, Path: "assets/images:assets/fonts"}: {
			metadataloader.Load,
		},
		{Embed: assets, Path: "assets/audio"}: {
			audioloader.Load,
		},
		{Embed: assets, Path: "assets/fonts"}: {
			fontloader.Load,
		},
	})

	printer.PrintCliMessage("HideSeek\nClient!")

	if cli.GetDebug(){
		profiling.UseProfiler().Init()
	}
	// collisions.ConnectCollisionsToImages()
}

func main() {
	ebiten.SetWindowSize(screen.GetMaxWidth(), screen.GetMaxHeight())
	ebiten.SetWindowTitle("HideSeek-Client")
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSizeLimits(screen.GetMinWidth(), screen.GetMinHeight(), -1, -1)

	

	log.Fatalln(ebiten.RunGame(loop.New()))
}
