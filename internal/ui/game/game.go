package GameProcess

import "github.com/hajimehoshi/ebiten/v2"

// Map "Game/Components/Map"
// "Game/Components/Sound"
// "Game/Components/States"
// "Game/Heroes/Animation"
// "Game/Heroes/Users"
// "Game/Server"
// "Game/UI/GameProcess/ConfigParsers"
// "Game/Window"

// "github.com/faiface/pixel"
// "github.com/faiface/pixel/pixelgl"

// type GameProcess struct {
// 	//It is such called stage struct
// 	//it uses all the important methods
// 	//for the corrisponding 'Stage' interface

// 	winConf *Window.WindowConfig

// 	currState *States.States

// 	userConfig *Users.User

// 	mapComponents Map.MapConf
// }

// func (g *GameProcess) Init(winConf *Window.WindowConfig, currState *States.States, userConfig *Users.User, mapComponents Map.MapConf) {
// 	g.winConf = winConf
// 	g.currState = currState
// 	g.userConfig = userConfig
// 	g.mapComponents = mapComponents
// }

// func GetUserFromList(u string, l []*Server.GameRequest) *Server.GameRequest {
// 	for _, value := range l {
// 		if value.PersonalInfo.Username == u {
// 			return value
// 		}
// 	}
// 	return nil
// }

// func (g *GameProcess) Run() {

// 	g.ProcessKeyboard()

// 	g.ProcessNetworking()

// 	g.DrawElements()

// 	g.ProcessMusic()
// }

func Draw(screen *ebiten.Image){
	// screen.DrawImage(, &ebiten.DrawImageOptions{})
	// g.winConf.DrawGameBackground()

	// // g.winConf.DrawGoldChest()

	// g.mapComponents.GetCollisions().GetDoorsCollisions().DrawDoors(g.winConf.DrawHorDoor, g.winConf.DrawVerDoor)

	// Animation.NewDefaultSwordAnimator(g.winConf, g.userConfig).Move()
	// Animation.NewIconAnimator(g.winConf, g.userConfig).Move()

	// for _, value := range g.winConf.GameProcess.OtherUsers {
	// 	Animation.NewDefaultSwordAnimator(g.winConf, value).Move()
	// 	Animation.NewIconAnimator(g.winConf, value).Move()
	// }

	// g.winConf.DrawDarkness(pixel.V((float64(g.userConfig.Pos.X)*2.5)-31, (float64(g.userConfig.Pos.Y)*2.5)-30))

	// g.winConf.DrawElementsPanel()


	// g.mapComponents.GetCam().UpdateCam()

	// var bias float64
	// for i := 0; i <= g.userConfig.GameInfo.Health; i++ {
	// 	g.winConf.DrawHPHeart(
	// 		pixel.V(-40+bias, 1200),
	// 	)
	// 	bias += 100
	// }

	// g.winConf.DrawWeaponIcon(g.userConfig.GameInfo.WeaponName)

	// if g.userConfig.GameInfo.Health < 1 {
	// 	g.mapComponents.GetCam().SetDefaultCam()
	// 	g.currState.MainStates.SetStartMenu()
	// }
}
