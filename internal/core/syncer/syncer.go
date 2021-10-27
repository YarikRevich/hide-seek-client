package syncer

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/events"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	screenhistory "github.com/YarikRevich/HideSeek-Client/internal/core/screen"
	"github.com/hajimehoshi/ebiten/v2"
)

type syncer struct {
	sw, sh, pw, ph float64
}

type SyncerProvider interface {
	UpdateScreenDeps(screen *ebiten.Image)
	Sync()
}

//Updates screen deps which syncer depends on
func (s *syncer) UpdateScreenDeps(screen *ebiten.Image){
	sw, sh := screen.Size()
	s.sw = float64(sw); s.sh = float64(sh)
	pw, ph := screenhistory.GetLastScreenSize()
	s.pw = float64(pw); s.ph = float64(ph)
}

//Sync data which depends on screen resize
func (s *syncer) Sync(){
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

			if v.Info.Scrollable {
				if v.RawMargins.TopMargin+e.MouseWheelY != v.Margins.TopMargin {
					v.Margins.TopMargin = v.RawMargins.TopMargin + e.MouseWheelY
				}
			}
		}
	}
}
	

func NewSyncer() SyncerProvider{
	return new(syncer)
}
