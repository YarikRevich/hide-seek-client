package objects

type Projectile struct {
	Base
	Opts ProjectileOpts
}

type ProjectileOpts struct {
	Angle float64
}

func NewProjectile(opts ProjectileOpts) *Projectile {
	return &Projectile{Opts: opts}
}
