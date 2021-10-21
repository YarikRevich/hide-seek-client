package position

import (
	"image"

	"github.com/google/uuid"
)

var positionCollection = make(map[uuid.UUID]image.Point)

func GetLastPosition(id uuid.UUID)image.Point{
	return positionCollection[id]
}

func SetLastPosition(id uuid.UUID, p image.Point){
	positionCollection[id] = p
}