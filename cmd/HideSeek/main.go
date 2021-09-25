package main

import (
	"flag"
	"log"
	"sync"

	"github.com/YarikRevich/HideSeek-Client/internal/loop"
	"github.com/YarikRevich/HideSeek-Client/internal/ai/collisions"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader"
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/image_loader"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/paths"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenWidth, screenHeight = InitScreenSize()
	fullWidth, fullHeight = ebiten.ScreenSizeInFullscreen()
)

func InitScreenSize()(int, int){
	return int(float64(fullWidth) / 1.15), int(float64(fullHeight) / 1.15)
}

func init(){
	flag.Parse()
	loader.LoadResources(map[string][]func(string, string, string, *sync.WaitGroup){
		paths.GAME_ASSETS_DIR: {
			imageloader.Load,
			metadataloader.Load,
		},
	})
	collisions.ConnectCollisionsToImages()
}

	//Gets info from user to place his name and server's name
	// username, server := Start.GetStartInfo()
	// conn := Server.GetConnection(server)
	// defer conn.Close()

	// //Create window and place all the components
	// winConf := Window.CreateWindow()

	// //Loads all the components
	// winConf.LoadAllImageComponents()
	// winConf.LoadAllTextAreas()
	// winConf.LoadAvailableHeroImages()
	// winConf.LoadAvailableWeaponImages()
	// winConf.LoadAvailableWeaponIconImages()

	// //Draws background image to start game
	// winConf.DrawBackgroundImage()

	//Get user's spawn place
	// randomSpawn := Utils.GetRandomSpawn()

	//Configures user's info


	//Sets current state at 'StartMenu'
	// currState := States.States{
	// 	MainStates:       new(States.MainStates),
	// 	MusicStates:      new(States.MusicStates),
	// 	SendStates:       new(States.SendStates),
	// 	NetworkingStates: new(States.NetworkingStates),
	// }
// currState.MainStates.SetStartMenu()

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
	ebiten.SetWindowSizeLimits(int((fullWidth * 60)/ 100), int((fullHeight * 60)/ 100), -1, -1)

	log.Fatalln(ebiten.RunGame(loop.New()))
}
