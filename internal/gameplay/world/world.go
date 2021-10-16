package world

import (
	"fmt"
	"strings"

	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/pc"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	metadatamodels "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

var instance *World

type Location struct {
	Name     string
	Path     string
	Image    *ebiten.Image `json:"-"`
	Metadata metadatamodels.Metadata
}

type World struct {
	Location

	ID uuid.UUID

	Users []pc.PC
}

func (w *World) Init(path string) {
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	w.ID = id

	w.Path = path
	split := strings.Split(path, "/")
	w.Name = split[len(split)-1]
	w.Image = imagecollection.GetImage(w.Path)
	w.Metadata = *metadatacollection.GetMetadata(w.Path)
}

func (w *World) Reset() {
	w.Users = w.Users[:0]
}

func (w *World) FormatUsersUsername() string {
	var r string
	for _, v := range w.Users {
		r += fmt.Sprintf("%s\n", v.Username)
	}
	return r
}

func UseWorld() *World {
	if instance == nil {
		instance = new(World)
	}
	return instance
}
