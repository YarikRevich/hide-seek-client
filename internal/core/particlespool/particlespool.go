package particlespool

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/engoengine/glm"
)

var instance *ParticlePool

const POOL_CAPACITY = 4

type Props struct {
	Position, Velocity, VelocityVariation [2]float64
	Life                                  float64
	ColorBegin, ColorEnd                  glm.Quat
	SizeBegin, SizeEnd, SizeVariation     glm.Quat
	LifeTime                              float64
}

type Particle struct {
	Position             [2]float64
	Velocity             [2]float64
	ColorBegin, ColorEnd glm.Quat

	Rotation           float64
	SizeBegin, SizeEnd glm.Quat
	LifeTime           float64
	LifeRemaining      float64
	Active             bool
}

type ParticlePool struct {
	pool  []*Particle
	index int

	particleAmount int
}

func (pm *ParticlePool) ForEachParticle(c func(*Particle)) {
	for _, v := range pm.pool {
		if v != nil && v.Active {
			c(v)
		}
	}
}

func (pm *ParticlePool) Fill(props Props) {
	for i := 0; i < len(pm.pool); i++ {
		fmt.Println(i)
		particle := pm.pool[i]
		if particle != nil && particle.Active {
			continue
		}

		if particle == nil {
			pm.pool[i] = new(Particle)
			particle = pm.pool[i]
		}

		particle.Active = true
		particle.Position = props.Position
		particle.Rotation = float64(rand.Intn(10)) * math.Pi

		particle.Velocity = props.Velocity
		particle.Velocity[0] += props.VelocityVariation[0] * float64(rand.Intn(10)) * math.Pi
		particle.Velocity[1] += props.VelocityVariation[1] * float64(rand.Intn(10)) * math.Pi

		particle.ColorBegin = props.ColorBegin
		particle.ColorEnd = props.ColorEnd

		particle.LifeTime = props.LifeTime
		particle.LifeRemaining = props.LifeTime

		svc := float32(rand.Intn(10)) - 0.5
		n := props.SizeVariation.Mul(&glm.Quat{V: glm.Vec3{svc, svc}})
		particle.SizeBegin = props.SizeBegin.Add(&n)
		particle.SizeEnd = props.SizeEnd
	}
}

func (pm *ParticlePool) Update(coef float64) {
	for _, v := range pm.pool {
		if v != nil {
			if !v.Active {
				continue
			}
			if v.LifeRemaining <= 0 {
				v.Active = false
				continue
			}

			v.LifeRemaining -= coef
			v.Position[0] += v.Velocity[0] * coef
			v.Position[1] += v.Velocity[1] * (coef / 12)
			v.Rotation += 0.01 * coef
		}
	}
}

func (pm *ParticlePool) Clean() {
	pm.pool = make([]*Particle, pm.particleAmount)
}
func Use() *ParticlePool {
	if instance == nil {
		instance = &ParticlePool{
			pool:           make([]*Particle, POOL_CAPACITY),
			index:          POOL_CAPACITY - 1,
			particleAmount: POOL_CAPACITY,
		}
	}
	return instance
}
