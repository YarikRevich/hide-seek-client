package pc

import (
	"crypto/sha256"
	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/YarikRevich/HideSeek-Client/internal/storage/provider"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	EMPTY          = ""
	DEFAULT_HEALTH = 10
)

var instance *PC

type Animation struct {
	FrameCount         uint32
	FrameDelay         uint32
	FrameDelayCounter  uint32
	CurrentFrameMatrix []float64
}

type GameCredentials struct {
	LobbyID string
}

type Skin struct {
	ImageHash [sha256.Size]byte
	Animation Animation
}

type Weapon struct {
	Name      string
	Radius    int
	Animation Animation
}

type Equipment struct {
	Skin   Skin
	Weapon string
}

type Buffs struct {
	Speed float64
}

type Physics struct {
	Jump []direction.Direction
}

type PC struct {
	ID uuid.UUID

	Username string

	Health uint64

	X float64
	Y float64

	Buffs Buffs

	Physics   Physics
	Equipment Equipment

	Metadata *models.Metadata

	GameCredentials GameCredentials
}

func (p *PC) Init(spawns []struct {
	X float64
	Y float64
}) {
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Fatal("failed to create uuid for world:", err)
	}
	p.ID = id

	n, ok := provider.UseStorageProvider().User().Get("name").(string)
	if !ok {
		logrus.Fatal("username can't be converted to string type")
	}
	p.Username = n

	x, y := GetSpawn(spawns)
	instance.X = x
	instance.Y = y
}

func UsePC() *PC {
	if instance == nil {
		instance = new(PC)
		instance.Username = EMPTY

		instance.Health = DEFAULT_HEALTH
		instance.Buffs.Speed = 2.5
		instance.Equipment.Skin.Animation.FrameDelay = 5
		instance.Equipment.Skin.Animation.FrameDelayCounter = 1
		instance.Metadata = metadatacollection.GetMetadata("assets/images/heroes/pumpkinhero")
	}
	return instance
}
