package collisions

import (
	"crypto/sha256"

	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/YarikRevich/HideSeek-Client/tools/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Collisions = make(map[[sha256.Size]byte]*models.Metadata)
)

func ConnectCollisionsToImages() {
	for hash := range imagecollection.ImageCollection {

		path := utils.GetPathByHash(hash, imagecollection.PathsToHash)
		coll, ok := metadatacollection.MetadataCollection[path]
		if ok {
			Collisions[hash] = coll
		}
	}
}

func SyncCollisionWithImage(screen *ebiten.Image, i *ebiten.Image) {
	//sets x and y of coll due to screen
	// id := i.Id.String()
	// Collisions[id]
}
