package syncer

import (
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
)

func SyncPC(swf, shf, pwf, phf float64){
	p := pc.UsePC()

	p.Y = (p.Y * shf) / phf
	p.X = (p.X * swf) / pwf

	p.Metadata.Size.Height = (p.Metadata.Size.Height * shf) / phf
	p.Metadata.Size.Width = (p.Metadata.Size.Width * swf) / pwf

	p.Buffs.Speed = (p.Buffs.Speed * swf) / pwf
}