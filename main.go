package main

import (
	"flag"
	"log"
	"os"

	"embed"

	"github.com/YarikRevich/HideSeek-Client/internal/ai/collisions"
	"github.com/YarikRevich/HideSeek-Client/internal/loop"
	"github.com/YarikRevich/HideSeek-Client/internal/paths"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/audio_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/font_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader"
	"github.com/YarikRevich/HideSeek-Client/tools/cli"
	"github.com/sirupsen/logrus"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed assets/images
	images embed.FS

	//go:embed assets/audio
	audio embed.FS

	//go:embed assets/fonts
	fonts embed.FS
)

var (
	screenWidth, screenHeight = InitScreenSize()
	fullWidth, fullHeight     = ebiten.ScreenSizeInFullscreen()
)

func InitScreenSize() (int, int) {
	return int(float64(fullWidth) / 1.15), int(float64(fullHeight) / 1.15)
}

func init() {
	flag.Parse()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	lgf, err := os.OpenFile(paths.GAME_LOG_DIR+"/log.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln(err)
	}
	logrus.SetOutput(lgf)

	if cli.GetDebug() {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetLevel(logrus.WarnLevel)

	resource_manager.LoadResources(map[resource_manager.Component][]resource_manager.Loader{
		{Embed: images, Path: "assets/images"}: {
			imageloader.Load,
			metadataloader.Load,
		},
		{Embed: audio, Path: "assets/audio"}: {
			audioloader.Load,
		},
		{Embed: fonts, Path: "assets/fonts"}: {
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
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hide&Seek")
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSizeLimits(int((fullWidth*60)/100), int((fullHeight*60)/100), -1, -1)

	log.Fatalln(ebiten.RunGame(loop.New()))
}
