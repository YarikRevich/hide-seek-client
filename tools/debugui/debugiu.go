package debugui

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/YarikRevich/hide-seek-client/internal/core/world"
	"github.com/YarikRevich/hide-seek-client/tools/debugui/scenes/game"
	"github.com/YarikRevich/hide-seek-client/tools/debugui/scenes/waitroomjoin"
	"github.com/YarikRevich/hide-seek-client/tools/params"
	"github.com/gabstv/ebiten-imgui/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/inkyblackness/imgui-go/v2"
)

var instance *DebugImGUI

type DebugImGUI struct {
	renderer *renderer.Manager
}

func (d *DebugImGUI) Update() {
	width, height := ebiten.WindowSize()
	d.renderer.Update(1.0/60.0, float32(width), float32(height))
}

func (d *DebugImGUI) Render(screen *ebiten.Image) {
	d.renderer.BeginFrame()
	pc := world.UseWorld().GetPC()
	{
		imgui.Begin("Debug menu")

		if imgui.BeginMenu("Actions") {
			if imgui.Button("Kill pc") {
				statemachine.UseStateMachine().PCs().SetState(pc.ID, statemachine.PC_DEAD_NOW)
			}

			imgui.EndMenu()
		}

		if imgui.BeginMenu("Scenes") {
			if imgui.Button("Game") {
				game.Show()
			}

			if imgui.Button("Waitroomjoin") {
				waitroomjoin.Show()
			}

			imgui.EndMenu()
		}

		if imgui.BeginMenu("Commands") {
			var (
				label  string
				action bool
			)
			if params.IsWithoutSound() {
				label = "Enable sound"
				action = false
			} else {
				label = "Disable sound"
				action = true
			}
			if imgui.Button(label) {
				params.SetWithoutSoundManually(action)
			}
			imgui.EndMenu()
		}

		imgui.End()
	}
	d.renderer.EndFrame(screen)
}

func UseDebugImGUI() *DebugImGUI {
	if instance == nil {
		instance = &DebugImGUI{renderer.New(nil)}
		imgui.CurrentIO().SetIniFilename("~/imgui.ini")
	}
	return instance
}
