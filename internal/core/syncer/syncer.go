package syncer

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

type syncer struct {
	sw, sh, pw, ph float64
}

type SyncerProvider interface {
	UpdateScreenDeps()
	Sync()
}

//Updates screen deps which syncer depends on
func (s *syncer) UpdateScreenDeps() {
	sw, sh := screenhistory.GetScreen().Size()
	s.sw = float64(sw)
	s.sh = float64(sh)
	pw, ph := screenhistory.GetLastScreenSize()
	s.pw = float64(pw)
	s.ph = float64(ph)
}

//Sync data which depends on screen resize
func (s *syncer) Sync() {
	if s.pw != 0 && s.ph != 0 {
		e := events.UseEvents().Mouse()
		for _, v := range metadatacollection.MetadataCollection {
			v.Scale.CoefficiantY = (v.Scale.CoefficiantY * s.sh) / s.ph
			v.Scale.CoefficiantX = (v.Scale.CoefficiantX * s.sw) / s.pw

			if v.RawSize.Width*v.Scale.CoefficiantX != v.Size.Width {
				v.Size.Width = v.RawSize.Width * v.Scale.CoefficiantX
			}

			if v.RawSize.Height*v.Scale.CoefficiantY != v.Size.Height {
				v.Size.Height = v.RawSize.Height * v.Scale.CoefficiantY
			}

			if v.Buffs.RawSpeed.X*v.Scale.CoefficiantX != v.Buffs.Speed.X {
				v.Buffs.Speed.X = v.Buffs.RawSpeed.X * v.Scale.CoefficiantX
			}

			if v.Buffs.RawSpeed.Y*v.Scale.CoefficiantY != v.Buffs.Speed.Y {
				v.Buffs.Speed.Y = v.Buffs.RawSpeed.Y * v.Scale.CoefficiantY
			}

			if v.Info.ScrollableX{
				if v.RawMargins.LeftMargin+e.MouseWheelX != v.Margins.LeftMargin {
					v.Margins.LeftMargin = v.RawMargins.LeftMargin + e.MouseWheelX
				}
			}

			if v.Info.ScrollableY {
				if v.RawMargins.TopMargin+e.MouseWheelY != v.Margins.TopMargin {
					v.Margins.TopMargin = v.RawMargins.TopMargin + e.MouseWheelY
				}
			}
		}
	}
}

func NewSyncer() SyncerProvider {
	return new(syncer)
}
