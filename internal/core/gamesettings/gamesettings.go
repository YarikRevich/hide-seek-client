package gamesettings

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/gamesettings/regime"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
)

type GameSettings struct {
	IsWorldExist, IsGameStarted bool
	Regime                      regime.Regime
}

func (gs *GameSettings) SetWorldExist(s bool) {
	gs.IsWorldExist = s
}

func (gs *GameSettings) SetGameStarted(s bool) {
	gs.IsGameStarted = s
}

func (gs *GameSettings) SetRegime(r regime.Regime) {
	gs.Regime = r
}

func (gs *GameSettings) FromAPIMessage() {

}

func (gs *GameSettings) ToAPIMessage() *api.GameSettings {
	return &api.GameSettings{
		IsGameStarted: gs.IsGameStarted,
		IsWorldExist:  gs.IsWorldExist,
	}
}

func NewGameSettings() *GameSettings {
	return new(GameSettings)
}
