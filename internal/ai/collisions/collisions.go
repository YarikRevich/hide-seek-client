package collisions

import (
	"crypto/sha256"

	metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"
	imageloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/image_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Collisions = make(map[[sha256.Size]byte]metadataloader.M)
)

func ConnectCollisionsToImages(){
	for path, img := range imageloader.Images{
		coll, ok := metadataloader.Metadata[path]
		if ok{
			Collisions[img.Id] = coll
		}
	}
}

func SyncCollisionWithImage(screen *ebiten.Image, i *imageloader.Image){
	//sets x and y of coll due to screen
	// id := i.Id.String()
	// Collisions[id]
}