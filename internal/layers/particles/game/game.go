package game

import (
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/particlespool"
	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/engoengine/glm"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	render.UseRender().SetToRender(func(i *ebiten.Image) {
		pool := particlespool.Use()
		if events.UseEvents().Mouse().IsAnyMovementButtonPressed() {

			particle := particlespool.Props{
				SizeBegin:         glm.Quat{V: glm.Vec3{0.5}},
				SizeVariation:     glm.Quat{V: glm.Vec3{0.3}},
				SizeEnd:           glm.Quat{V: glm.Vec3{0.0}},
				LifeTime:          1.0,
				Velocity:          [2]float64{0.0, 0.0},
				VelocityVariation: [2]float64{3.0, 1.0},
				Position:          [2]float64{50.0, 50.0},
				ColorBegin:        glm.Quat{W: 0, V: glm.Vec3{255, 255, 243}},
				ColorEnd:          glm.Quat{W: 0, V: glm.Vec3{255, 234, 243}},
			}

			pool.Fill(particle)
		}

		pool.ForEachParticle(func(p *particlespool.Particle) {
			particleImage := i.SubImage(image.Rect(0, 10, 10, 10)).(*ebiten.Image)

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(p.Position[0], p.Position[1])
			opts.GeoM.Rotate(p.Rotation)

			life := float32(p.LifeRemaining / p.LifeTime)

			color := glm.QuatLerp(&p.ColorEnd, &p.ColorBegin, life)
			opts.ColorM.Translate(float64(color.X()), float64(color.Y()), float64(color.Z()), float64(color.W))

			scale := glm.QuatLerp(&p.SizeEnd, &p.SizeBegin, life)
			opts.GeoM.Scale(float64(scale.X()), float64(scale.X()))

			i.DrawImage(particleImage, opts)
		})
	})
}
