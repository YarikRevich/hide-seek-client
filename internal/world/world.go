package world

import metadataloader "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/loader/metadata_loader"

var (
	w *World
)

type World struct {
	Map string

	Metadata metadataloader.Metadata
}

func UseWorld()*World{
	if w == nil{
		w = new(World)
	}
	return w
}