package collisions

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/objects"
	"github.com/YarikRevich/hide-seek-client/internal/core/sources"
)

var instance *CollisionDetector

type CollisionDetector struct {
}

//Returns information about collision
func (cd *CollisionDetector) GetObjectsCollision(v1, v2 *objects.Base) (*sources.CollidersModel, bool) {
	// v1Size := v1.GetSize()
	// v2Size := v2.GetSize()
	// c := camera.UseCamera()
	// func (o *Base) GetRect() image.Rectangle {
	// 	s := o.GetSize()
	// 	c := camera.UseCamera()

	// 	return image.Rect(int(o.RawPos.X), int(o.RawPos.Y), int(o.RawPos.X+s.X), int(o.RawPos.Y+s.Y))
	// 	// // func (m *MetadataModel) GetRect() image.Rectangle {
	// 	// 	ms := m.GetMargins()

	// 	// 	ma := m.GetMargins()
	// 	// 	return image.Rect(int(ma.X), int(ma.Y), int(ms.X+s.X), int(ms.Y+s.Y))
	// 	// }
	// }

	return nil, false
}

//Returns information about collision
func (cd *CollisionDetector) GetTMXCollision(v1 *objects.Base, collisionBatch sources.CollidersBatch) (sources.CollidersModel, bool) {
	v1Size := v1.GetSize()
	// v1Margins := v1.GetMargins()
	c := camera.UseCamera()
	r := image.Rect(v1.RawPos.X, v1.RawPos.Y, v1.RawPos.X + )
	for q, v := range collisionBatch {
		if q.Min.X <= r.Max.X &&
			q.Max.X >= r.Min.X &&
			q.Min.Y <= r.Max.Y &&
			q.Max.Y >= r.Min.Y {
			return v, true
		}
	}

	return sources.CollidersModel{}, false
}

// sources.CollidersModel

func UseCollisionDetector() *CollisionDetector {
	if instance == nil {
		instance = &CollisionDetector{}
	}
	return instance
}
