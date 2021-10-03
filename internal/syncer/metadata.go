package syncer

import (
	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
)

func SyncMetadata(swf, shf, pwf, phf float64) {
	for _, v := range metadataloader.MetadataCollection {
		v.Scale.CoefficiantY = (v.Scale.CoefficiantY * shf) / phf
		v.Scale.CoefficiantX = (v.Scale.CoefficiantX * swf) / pwf
	}
}
