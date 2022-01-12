package collisions

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
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

var instance *Collisions

type Collisions struct {
	//Cache for performace
	cache map[*objects.Base][]*sources.CollidersModel
}

//Checks whether passed object collides with any position
//in a collider set
func (c *Collisions) IsCollide(object *objects.Base, collidersSet string) bool {
	cs := sources.UseSources().Colliders().GetCollider(collidersSet)
	for _, v := range c.cache {
		for _, e := range v {
			for _, q := range cs {
				if e.X == q.X && e.Y == q.Y {
					return true
				}
			}
		}
	}

	for _, q := range cs {

		//Checks for collision detection
		if ((object.ModelCombination.Modified.GetSizeMinX() <= q.GetSizeMaxX() &&
			object.ModelCombination.Modified.GetSizeMinX() >= q.GetSizeMinX()) || (object.ModelCombination.Modified.GetSizeMinX() >= q.GetSizeMaxX() &&
			object.ModelCombination.Modified.GetSizeMinX() <= q.GetSizeMinX())) &&
			((object.ModelCombination.Modified.GetSizeMinY() <= q.GetSizeMaxY() &&
				object.ModelCombination.Modified.GetSizeMinY() >= q.GetSizeMinY()) || (object.ModelCombination.Modified.GetSizeMinY() >= q.GetSizeMaxY() &&
				object.ModelCombination.Modified.GetSizeMinY() <= q.GetSizeMinY())) {
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

func UseCollisions() *Collisions {
	if instance == nil {
		instance = &Collisions{
			cache: make(map[*objects.Base][]*sources.CollidersModel),
		}
	}
	return instance
}
