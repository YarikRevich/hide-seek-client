package sources

import (
	"strings"

	"github.com/YarikRevich/hide-seek-client/assets"
	"github.com/YarikRevich/hide-seek-client/internal/core/screen"
	"github.com/hajimehoshi/ebiten/v2"
)

type Shader struct {
	Name   string
	Shader *ebiten.Shader
}

func (s *Shader) load(path string) error {
	file, err := assets.Assets.ReadFile(path)
	if err != nil {
		return err
	}

	shader, err := ebiten.NewShader(file)
	if err != nil {
		return err
	}
	s.Shader = shader

	s.Name = strings.Split(path, ".")[0]
	shaderCollection[s.Name] = s

	return nil
}

func (s *Shader) OnInteract() {

}

func (s *Shader) OnCollision() {

}

func (s *Shader) Render(sm screen.ScreenManager) {

}
