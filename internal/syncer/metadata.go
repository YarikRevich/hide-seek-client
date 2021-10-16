package syncer

import metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"

func SyncMetadata(swf, shf, pwf, phf float64) {
	for _, v := range metadatacollection.MetadataCollection {
		v.Scale.CoefficiantY = (v.Scale.CoefficiantY * shf) / phf
		v.Scale.CoefficiantX = (v.Scale.CoefficiantX * swf) / pwf

		v.Size.Width = (v.Size.Width * shf) / phf
		v.Size.Height = (v.Size.Height * swf) / pwf

		v.Margins.LeftMargin = (v.Margins.LeftMargin * shf) / phf
		v.Margins.TopMargin = (v.Margins.TopMargin * swf) / pwf
	}
}
