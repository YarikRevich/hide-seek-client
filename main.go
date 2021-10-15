package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"embed"

	"github.com/YarikRevich/HideSeek-Client/internal/ai/collisions"
	"github.com/YarikRevich/HideSeek-Client/internal/loop"
	"github.com/YarikRevich/HideSeek-Client/internal/paths"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/screen"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
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

	lgf, err := os.OpenFile(filepath.Join(paths.GAME_LOG_DIR, "/log.log"), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln(err)
	}
	logrus.SetOutput(lgf)

	if cli.GetDebug() {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetLevel(logrus.WarnLevel)

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
	collisions.ConnectCollisionsToImages()
}

//Get user's spawn place
// randomSpawn := Utils.GetRandomSpawn()

//Configures map
// mapComponents := Map.MapConf(new(Map.MapC))
// mapComponents.Init()
// mapComponents.ConfAll(&winConf, userConfig)

// winConf.Components.AvPlacesForSpaws = mapComponents.GetAnailizer().AnalizeAvailablePlaces()

//Starts pinger
// go log.GetPing()

// for !winConf.Win.Closed() {

// 	//fmt.Println(currState.MainStates)
// 	//Shows statistics about user if argument is placed
// 	log.Show()

// 	frames++
// 	select {
// 	case <-second:
// 		//Sets title of the window with frame rate
// 		winConf.Win.SetTitle(fmt.Sprintf("Hide and seek| %d", frames))
// 		frames = 0
// 	default:
// 		//Upgrades background
// 		winConf.UpdateBackground()

// 		//Goes to the action gate to chose an important one
// 		choseActionGate(&winConf, &currState, &userConfig, mapComponents)

// 		winConf.Win.Update()
// 	}
// }
// color.Green.Println("Goodbye!")
// }

func main() {
	ebiten.SetWindowSize(screen.GetMaxWidth(), screen.GetMaxHeight())
	ebiten.SetWindowTitle("Hide&Seek")
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSizeLimits(screen.GetMinWidth(), screen.GetMinHeight(), -1, -1)

	log.Fatalln(ebiten.RunGame(loop.New()))
}
