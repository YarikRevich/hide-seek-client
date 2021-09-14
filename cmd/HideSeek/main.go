package main

//go:generate sh ../scripts/init.sh
import (
	
	// "github.com/faiface/pixel/pixelgl"
	// "github.com/gookit/color"

	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/game"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/image_loader"
	"github.com/YarikRevich/Hide-Seek-with-Guns/internal/messages"
	"github.com/hajimehoshi/ebiten/v2"
)

func init(){
	messages.GetInstance().Listen()
	imageloader.LoadImages()
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
	fullWidth, fullHeight := ebiten.ScreenSizeInFullscreen()
	
	ebiten.SetWindowSize(int(float64(fullWidth) / 1.25), int(float64(fullHeight) / 1.25))
	ebiten.SetWindowTitle("Hide&Seek")

	ebiten.RunGame(game.New())
}
