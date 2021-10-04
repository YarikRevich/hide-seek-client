package syncer

import (
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

func SyncMetadata(swf, shf, pwf, phf float64) {
	for _, v := range metadatacollection.MetadataCollection {
		v.Scale.CoefficiantY = (v.Scale.CoefficiantY * shf) / phf
		v.Scale.CoefficiantX = (v.Scale.CoefficiantX * swf) / pwf
	}
}
