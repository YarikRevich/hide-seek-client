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
	for hash := range imageloader.Images{
		
		path := imageloader.GetPathByHash(hash)
		coll, ok := metadataloader.Metadata[path]
		if ok{
			Collisions[hash] = coll
		}
	}
}

func SyncCollisionWithImage(screen *ebiten.Image, i *ebiten.Image){
	//sets x and y of coll due to screen
	// id := i.Id.String()
	// Collisions[id]
}