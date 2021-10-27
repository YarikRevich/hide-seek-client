package objects

import "image"

//Returns spawn points for elements on the map
//there should be stated spawn zones on the map metadata
func GetSpawn(spawnPoints []image.Point) (float64, float64) {
	return 500, 500
}