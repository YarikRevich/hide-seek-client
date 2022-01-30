package collisions

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
	"github.com/YarikRevich/hide-seek-client/internal/core/types"
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

var instance *CollisionDetector

type CollisionDetector struct {
}

//Returns information about collision
func (cd *CollisionDetector) GetObjectsCollision(v1, v2 *objects.Base) (*sources.CollidersModel, bool) {
	return nil, false
}

//Returns information about collision
func (cd *CollisionDetector) GetTMXCollision(c types.Vec2) (*sources.CollidersModel, bool) {
	return nil, false
}

// sources.CollidersModel

func UseCollisionDetector() *CollisionDetector {
	if instance == nil {
		instance = &CollisionDetector{}
	}
	return instance
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
