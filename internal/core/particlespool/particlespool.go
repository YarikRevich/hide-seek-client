package particlespool

import (
	"math"
	"math/rand"

	"github.com/engoengine/glm"
)

var instance *ParticlePool

const POOL_CAPACITY = 5

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
		c(v)
	}
}

func (pm *ParticlePool) Fill(props Props) {
	for i := 0; i < len(pm.pool); i++ {
		particle := pm.pool[i]

		particle.Active = true
		particle.Position = props.Position
		particle.Rotation = float64(rand.Intn(2)) * math.Pi

		particle.Velocity = props.Velocity
		particle.Velocity[0] += props.VelocityVariation[0] * float64(rand.Intn(2)) * math.Pi
		particle.Velocity[1] += props.VelocityVariation[1] * float64(rand.Intn(2)) * math.Pi

		particle.ColorBegin = props.ColorBegin
		particle.ColorEnd = props.ColorEnd

		particle.LifeTime = props.LifeTime
		particle.LifeRemaining = props.LifeTime

		svc := float32(rand.Intn(2)) - 0.5
		n := props.SizeVariation.Mul(&glm.Quat{V: glm.Vec3{svc, svc}})
		particle.SizeBegin = props.SizeBegin.Add(&n)
		particle.SizeEnd = props.SizeEnd
	}
}

func (pm *ParticlePool) Update() {

	for _, v := range pm.pool {
		if !v.Active {
			continue
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
