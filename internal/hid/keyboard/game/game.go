package game

import (
	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	"github.com/YarikRevich/HideSeek-Client/internal/physics/jump"
	"github.com/hajimehoshi/ebiten/v2"
)

func Exec(){
	//currX, currY := ebiten.CursorPosition()
	//mapName := "/"

	// && !boarders.IsBoarder()
	if ebiten.IsKeyPressed(ebiten.KeyW){
		history.SetDirection(direction.UP)
		pc.GetPC().Y -= pc.GetPC().Buffs.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS){
		history.SetDirection(direction.DOWN)
		pc.GetPC().Y += pc.GetPC().Buffs.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD){
		history.SetDirection(direction.RIGHT)
		pc.GetPC().X += pc.GetPC().Buffs.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA){
		history.SetDirection(direction.LEFT)
		pc.GetPC().X -= pc.GetPC().Buffs.Speed
	}
	
	if ebiten.IsKeyPressed(ebiten.KeySpace){
		jump.CalculateJump()
		// pc.GetPC().X -= pc.GetPC().Buffs.Speed
	}


// 	currPosition := pixel.V(float64(g.userConfig.Pos.X), float64(g.userConfig.Pos.Y))
// 	g.mapComponents.GetCollisions().GetDoorsCollisions().DoorTraker(currPosition)

// 	switch {
// 	case ((g.winConf.Win.Pressed(pixelgl.KeyW) ||
// 		g.winConf.Win.Pressed(pixelgl.KeyA) ||
// 		g.winConf.Win.Pressed(pixelgl.KeyS) ||
// 		g.winConf.Win.Pressed(pixelgl.KeyD)) && g.winConf.Win.JustPressed(pixelgl.KeySpace)) ||
// 		g.winConf.Win.JustPressed(pixelgl.KeySpace):
// 		ok, user := g.mapComponents.GetCollisions().GetHeroCollisions().IsHero(currPosition, g.winConf.GameProcess.OtherUsers).Near(30, 37)
// 		if ok {
// 			g.userConfig.Context.Additional = append(g.userConfig.Context.Additional, user, "1")
// 			parser := Server.GameParser(new(Server.GameRequest))
// 			server := Server.Network(new(Server.N))
// 			server.Init(nil, g.userConfig, 1, nil, parser.Parse, "UpdateUsersHealth")

// 			server.Write()
// 			server.ReadGame(parser.Unparse)

// 			g.userConfig.Context.Additional = []string{}
// 		}
// 		fallthrough
// 	default:
// 		switch {
// 		case g.winConf.Win.Pressed(pixelgl.KeyW):
// 			coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.Pos.X), float64(g.userConfig.Pos.Y+2)), g.winConf.GameProcess.OtherUsers, "top")
// 			if coll {
// 				return
// 			}

// 			if g.userConfig.Pos.Y <= g.mapComponents.GetHeroBorder().Top() {
// 				g.userConfig.Pos.Y += 3
// 			}
// 			if g.winConf.Cam.CamPos.Y < g.mapComponents.GetCamBorder().Top()+50 {
// 				if g.userConfig.Pos.Y >= int(g.winConf.Win.Bounds().Center().Y) {
// 					g.winConf.Cam.CamPos.Y += 5
// 				}
// 			}
// 		case g.winConf.Win.Pressed(pixelgl.KeyA):
// 			coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.Pos.X-2), float64(g.userConfig.Pos.Y)), g.winConf.GameProcess.OtherUsers, "left")
// 			if coll {
// 				return
// 			}

// 			if g.userConfig.Pos.X >= g.mapComponents.GetHeroBorder().Left() {
// 				g.userConfig.Pos.X -= 3
// 			}
// 			if g.winConf.Cam.CamPos.X >= g.mapComponents.GetCamBorder().Left() {
// 				if g.userConfig.Pos.X <= int(g.winConf.Win.Bounds().Center().X) {
// 					g.winConf.Cam.CamPos.X -= 5
// 				}
// 			}
// 		case g.winConf.Win.Pressed(pixelgl.KeyS):
// 			coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.Pos.X), float64(g.userConfig.Pos.Y-2)), g.winConf.GameProcess.OtherUsers, "bottom")
// 			if coll {
// 				return
// 			}

// 			if g.userConfig.Pos.Y >= g.mapComponents.GetHeroBorder().Bottom() {
// 				g.userConfig.Pos.Y -= 3
// 			}
// 			if g.winConf.Cam.CamPos.Y >= g.mapComponents.GetCamBorder().Bottom() {
// 				if g.userConfig.Pos.Y <= int(g.winConf.Win.Bounds().Center().Y) {
// 					g.winConf.Cam.CamPos.Y -= 5
// 				}
// 			}
// 		case g.winConf.Win.Pressed(pixelgl.KeyD):
// 			coll := g.mapComponents.GetCollisions().IsCritical(pixel.V(float64(g.userConfig.Pos.X+2), float64(g.userConfig.Pos.Y)), g.winConf.GameProcess.OtherUsers, "right")
// 			if coll {
// 				return
// 			}

// 			if g.userConfig.Pos.X <= g.mapComponents.GetHeroBorder().Right() {
// 				g.userConfig.Pos.X += 3
// 			}
// 			if g.winConf.Cam.CamPos.X <= g.mapComponents.GetCamBorder().Right() {
// 				if g.userConfig.Pos.X >= int(g.winConf.Win.Bounds().Center().X) {
// 					g.winConf.Cam.CamPos.X += 5
// 				}
// 			}
// 		}
// 	}
// }
}