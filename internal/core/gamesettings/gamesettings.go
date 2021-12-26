package gamesettings

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/gamesettings/regime"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
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

func (gs *GameSettings) ToAPIMessage() *server_external.GameSettings {
	return &server_external.GameSettings{
		IsGameStarted: gs.IsGameStarted,
		IsWorldExist:  gs.IsWorldExist,
	}
}

func NewGameSettings() *GameSettings {
	return new(GameSettings)
}
