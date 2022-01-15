package collisions

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/sirupsen/logrus"
	// "github.com/YarikRevich/hide-seek-client/internal/core/objects"
	// "github.com/YarikRevich/hide-seek-client/internal/core/world"
)

// //Model for elements in collisions pool
// type CollisionPoolCell struct {
// 	ObjectPosition, Collider struct {
// 		X, Y float64
// 	}
// 	IsColliding bool
// }

// var instance *Collisions

// type Collisions struct {
// 	//Cache for performace
// 	cache map[*objects.Base][]*sources.CollidersModel
// }

//Checks whether passed object collides with any position
//in a collider set

func IsCollision(object *objects.Base, collidersSet string) bool {
	cs, err := sources.UseSources().Colliders().GetCollider(collidersSet)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, q := range cs {

		//Checks for collision detection
		if ((object.MetadataModel.GetSizeMinX() <= q.GetSizeMaxX() &&
			object.MetadataModel.GetSizeMinX() >= q.GetSizeMinX()) || (object.MetadataModel.GetSizeMinX() >= q.GetSizeMaxX() &&
			object.MetadataModel.GetSizeMinX() <= q.GetSizeMinX())) &&
			((object.MetadataModel.GetSizeMinY() <= q.GetSizeMaxY() &&
				object.MetadataModel.GetSizeMinY() >= q.GetSizeMinY()) || (object.MetadataModel.GetSizeMinY() >= q.GetSizeMaxY() &&
				object.MetadataModel.GetSizeMinY() <= q.GetSizeMinY())) {
			return true
		}
	}
	return false
}

// //Cleans cache checking if such collision still exists
// //in tiled map, because due to the resizing of the screen
// //some collisions may change
// func (c *Collisions) CleanCache() {
// 	cs := sources.UseSources().Colliders().Collection
// 	for k, v := range c.cache {
// 	colliderscache:
// 		for i, e := range v {
// 			for p, q := range cs {
// 				if p == e.Name {
// 					for _, r := range q {
// 						if e.X == r.X && e.Y == r.Y {
// 							continue colliderscache
// 						}
// 					}
// 				}
// 			}
// 			c.cache[k] = append(c.cache[k][:i], c.cache[k][i+1:]...)
// 		}
// 	}
// }

// func UseCollisions() *Collisions {
// 	if instance == nil {
// 		instance = &Collisions{
// 			cache: make(map[*objects.Base][]*sources.CollidersModel),
// 		}
// 	}
// 	return instance
// }
