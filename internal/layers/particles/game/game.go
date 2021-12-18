package game

import (
	"fmt"

	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/particlespool"
	"github.com/YarikRevich/HideSeek-Client/internal/core/primitives"
	"github.com/YarikRevich/HideSeek-Client/internal/core/render"
	"github.com/YarikRevich/HideSeek-Client/internal/core/world"
	"github.com/YarikRevich/HideSeek-Client/tools/color"
	"github.com/engoengine/glm"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw() {
	render.UseRender().SetToRender(func(i *ebiten.Image) {
		cam := world.UseWorld().GetCamera()
		pc := world.UseWorld().GetPC()

		pool := particlespool.Use()
		if events.UseEvents().Mouse().IsAnyMovementButtonPressed() {
			particle := particlespool.Props{
				SizeBegin:         glm.Quat{V: glm.Vec3{0.3}},
				SizeVariation:     glm.Quat{V: glm.Vec3{0.1}},
				SizeEnd:           glm.Quat{V: glm.Vec3{0.2}},
				LifeTime:          2.0,
				Velocity:          [2]float64{0.01, 0.01},
				VelocityVariation: [2]float64{2.0, 0.5},
				Position:          [2]float64{0.0, 30.0},
				ColorBegin:        color.CreateColorFromArray(pc.ModelCombination.Modified.Effects.TraceColorBegin),
				ColorEnd:          color.CreateColorFromArray(pc.ModelCombination.Modified.Effects.TraceColorEnd),
			}

			pool.Fill(particle)
		}

		pool.Update(0.2)

		pool.ForEachParticle(func(p *particlespool.Particle) {
			img := primitives.CreateSquare(20)
			opts := &ebiten.DrawImageOptions{}

			life := float32(p.LifeRemaining / p.LifeTime)
			scale := glm.QuatLerp(&p.SizeEnd, &p.SizeBegin, life)
			opts.GeoM.Scale(float64(scale.X()), float64(scale.X()))

			if pc.IsDirectionLEFT() {
				opts.GeoM.Rotate(p.Rotation)
				opts.GeoM.Translate((p.Position[0])+(pc.GetScaledOffsetX()-cam.AlignOffset.X), (p.Position[1])+(pc.GetScaledOffsetY()-cam.AlignOffset.Y))

			} else {
				opts.GeoM.Rotate(-p.Rotation)
				opts.GeoM.Translate(((pc.GetScaledOffsetX() - cam.AlignOffset.X) - (p.Position[0])), (p.Position[1])+(pc.GetScaledOffsetY()-cam.AlignOffset.Y))
			}

			colorVariantion := glm.QuatLerp(&p.ColorEnd, &p.ColorBegin, life)
			fmt.Println(color.CreateRGBAFromQuatColor(colorVariantion))
			img.Fill(color.CreateRGBAFromQuatColor(colorVariantion))

			i.DrawImage(img, opts)
		})
	})
}
