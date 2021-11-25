package gamesettings

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/gamesettings/regime"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
)

type GameSettings struct {
	IsGameStarted bool
	Regime        regime.Regime
}

func (gs *GameSettings) SetRegime(r regime.Regime) {
	gs.Regime = r
}

func (gs *GameSettings) FromAPIMessage() {

}

func (gs *GameSettings) ToAPIMessage() *api.GameSettings {
	return &api.GameSettings{}
}

func NewGameSettings() *GameSettings {
	return new(GameSettings)
}
