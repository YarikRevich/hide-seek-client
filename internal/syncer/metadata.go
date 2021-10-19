package syncer

import (
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

func SyncMetadata(swf, shf, pwf, phf float64) {
	for _, v := range metadatacollection.MetadataCollection {
		v.Scale.CoefficiantY = (v.Scale.CoefficiantY * shf) / phf
		v.Scale.CoefficiantX = (v.Scale.CoefficiantX * swf) / pwf

		if v.RawSize.Width * v.Scale.CoefficiantX != v.Size.Width{
			v.Size.Width = v.RawSize.Width * v.Scale.CoefficiantX
		}

		if v.RawSize.Height * v.Scale.CoefficiantY != v.Size.Height{
			v.Size.Height = v.RawSize.Height * v.Scale.CoefficiantY
		}
	}
}
