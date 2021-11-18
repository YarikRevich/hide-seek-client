package syncer

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/core/sources"
)

type syncer struct {
	sw, sh, pw, ph float64
}

type SyncerProvider interface {
	Sync()
}

//Sync data which depends on screen resize
func (s *syncer) Sync() {
	if s.pw != 0 && s.ph != 0 {
		for _, v := range sources.UseSources().Metadata().Collection {
			v.Modified.Scale.CoefficiantX = (v.Modified.Scale.CoefficiantX * s.sw) / s.pw
			v.Modified.Scale.CoefficiantY = (v.Modified.Scale.CoefficiantY * s.sh) / s.ph

			if v.Origin.Size.Width*v.Modified.Scale.CoefficiantX != v.Modified.Size.Width {
				v.Modified.Size.Width = v.Origin.Size.Width * v.Modified.Scale.CoefficiantX
			}

			if v.Origin.Size.Height*v.Modified.Scale.CoefficiantY != v.Modified.Size.Height {
				v.Modified.Size.Height = v.Origin.Size.Height * v.Modified.Scale.CoefficiantY
			}

			if v.Origin.Buffs.Speed.X*v.Modified.Scale.CoefficiantX != v.Modified.Buffs.Speed.X {
				v.Modified.Buffs.Speed.X = v.Origin.Buffs.Speed.X * v.Modified.Scale.CoefficiantX
			}

			if v.Origin.Buffs.Speed.Y*v.Modified.Scale.CoefficiantY != v.Modified.Buffs.Speed.Y {
				v.Modified.Buffs.Speed.Y = v.Origin.Buffs.Speed.Y * v.Modified.Scale.CoefficiantY

			}

			e := events.UseEvents().Mouse()
			if v.Origin.Info.ScrollableX {
				if v.Origin.Margins.LeftMargin+e.MouseWheelX != v.Modified.Margins.LeftMargin {
					v.Modified.Margins.LeftMargin = v.Origin.Margins.LeftMargin + e.MouseWheelX
				}
			}

			if v.Origin.Info.ScrollableY {
				if v.Origin.Margins.TopMargin+e.MouseWheelY != v.Modified.Margins.TopMargin {
					v.Modified.Margins.TopMargin = v.Origin.Margins.TopMargin + e.MouseWheelY
				}
			}
		}
	}
}

func NewSyncer() SyncerProvider {
	s := new(syncer)
	sw, sh := screenhistory.UseScreen().GetScreen().Size()
	s.sw = float64(sw)
	s.sh = float64(sh)
	pw, ph := screenhistory.UseScreen().GetLastSize()
	s.pw = float64(pw)
	s.ph = float64(ph)
	return s
}
