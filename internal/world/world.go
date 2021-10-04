package world

import (
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
)

var (
	w *World
)

type World struct {
	Map string

	Metadata models.Metadata
}

func UseWorld()*World{
	if w == nil{
		w = new(World)
		w.Metadata = *metadatacollection.GetMetadata("assets/images/maps/default/background/Game")
	}
	return w
}