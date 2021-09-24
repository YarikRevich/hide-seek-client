package collisions

import (
	collisionloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/collision_loader"
	imageloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/image_loader"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Collisions = make(map[*ebiten.Image]collisionloader.Collision)
)

func ConnectCollisionsToImages(){
	for path, img := range imageloader.Images{
		coll, ok := collisionloader.RawCollisions[path]
		if ok{
			Collisions[img] = coll
		}
	}
}

// func SyncCollision(screen *ebiten.Image, coll *Collision){
// 	//sets x and y of coll due to screen
// }