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
