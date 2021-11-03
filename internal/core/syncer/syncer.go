package syncer

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	"github.com/YarikRevich/HideSeek-Client/internal/core/objects"
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

type syncer struct {
	sw, sh, pw, ph float64
}

type SyncerProvider interface {
	Sync()
}

//Syncs scale of pc if window is resized
func (s *syncer) syncScale() {
	p := objects.UseObjects().PC()
	p.Scale.CoefficiantX = (p.Scale.CoefficiantX * s.sw) / s.pw
	p.Scale.CoefficiantY = (p.Scale.CoefficiantY * s.sh) / s.ph
}

//Syncs margins of pc if window is resized
func (s *syncer) syncSize() {
	p := objects.UseObjects().PC()
	m := p.GetMetadata()

	if m.Size.Width*p.Scale.CoefficiantX != p.Size.Width {
		p.Size.Width = m.Size.Width * p.Scale.CoefficiantX
	}

	if m.Size.Height*p.Scale.CoefficiantY != p.Size.Height {
		p.Size.Height = m.Size.Height * p.Scale.CoefficiantY
	}
}

//Syncs margins of pc if window is resized
func (s *syncer) syncSpeed() {
	p := objects.UseObjects().PC()
	m := p.GetMetadata()
	
	if m.Buffs.Speed.X*p.Scale.CoefficiantX != p.Buffs.Speed.X {
		p.Buffs.Speed.X = m.Buffs.Speed.X * p.Scale.CoefficiantX
	}

	if m.Buffs.Speed.Y*p.Scale.CoefficiantY != p.Buffs.Speed.Y {
		p.Buffs.Speed.Y = m.Buffs.Speed.Y * p.Scale.CoefficiantY
	}
}

//Syncs margins of pc if window is resized
func (s *syncer) syncMarginsX() {
	p := objects.UseObjects().PC()
	m := p.GetMetadata()
	e := events.UseEvents().Mouse()
	if m.Margins.LeftMargin+e.MouseWheelX != p.Margins.LeftMargin {
		p.Margins.LeftMargin = m.Margins.LeftMargin + e.MouseWheelX
	}
}

//Syncs margins of pc if window is resized
func (s *syncer) syncMarginsY() {
	p := objects.UseObjects().PC()
	m := p.GetMetadata()
	e := events.UseEvents().Mouse()

	if m.Margins.TopMargin+e.MouseWheelY != p.Margins.TopMargin {
		p.Margins.TopMargin = m.Margins.TopMargin + e.MouseWheelY
	}
}

//Sync data which depends on screen resize
func (s *syncer) Sync() {
	if s.pw != 0 && s.ph != 0 {
		w := objects.UseObjects().World()
		// for _, v := range metadatacollection.MetadataCollection {
		for _, v := range w.Ammo{
			s.syncScale()
			s.syncSize()
			s.syncSpeed()

			if v.Info.ScrollableX {
				s.syncMarginsX()
			}

			if v.Info.ScrollableY {
				s.syncMarginsY()
			}
		// }
	}
}

func NewSyncer() SyncerProvider {
	s := new(syncer)
	sw, sh := screenhistory.GetScreen().Size()
	s.sw = float64(sw)
	s.sh = float64(sh)
	pw, ph := screenhistory.GetLastScreenSize()
	s.pw = float64(pw)
	s.ph = float64(ph)
	return s
}
