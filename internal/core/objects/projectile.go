package objects

type Projectile struct {
	Base
	Opts ProjectileOpts
}

type ProjectileOpts struct {
	Angle, Lifetime float64
}

func NewProjectile(opts ProjectileOpts) *Projectile {
	return &Projectile{Opts: opts}
}
