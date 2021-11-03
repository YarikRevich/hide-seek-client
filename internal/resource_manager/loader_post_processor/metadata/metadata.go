package metadata

import (
	"path/filepath"

	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
)

//Postprocesor which connects loaded
//image sizes to the proper metadata entry
func ConnectAdditionalStatementsToMetadata() {
	for k, v := range metadatacollection.MetadataCollection {
		if len(v.Info.Parent) != 0 {
			dir, _ := filepath.Split(k)
			imageW, imageH := imagecollection.GetImage(filepath.Join(dir, v.Info.Parent)).Size()
			if v.Animation.FrameNum != 0 {
				imageW /= int(v.Animation.FrameNum)
			}
			v.Size.Width = float64(imageW)
			v.Size.Height = float64(imageH)
		}
	}
}

//Postprocessor which connects metadata info
//to collision collection
func ConnectMetadataToCollisions() {}
