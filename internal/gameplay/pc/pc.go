package pc

import (
	"crypto/sha256"
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/models"
	"github.com/YarikRevich/HideSeek-Client/internal/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/storage/provider"
	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
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
	SpeedX float64
	SpeedY float64
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

	PositionHistory zeroshifter.IZeroShifter `json:"-"`

	Buffs Buffs

	Physics   Physics
	Equipment Equipment

	Image    *ebiten.Image `json:"-"`
	Metadata *models.Metadata

	GameCredentials GameCredentials
}

//Initializes pc username by requesting storage
func (p *PC) InitUsername() {
	n, ok := provider.UseStorageProvider().User().Get("name").(string)
	if !ok {
		logrus.Fatal("username can't be converted to string type")
	}
	p.Username = n
}

//Sets pc spawn by using configured map spawns
func (p *PC) SetSpawn(spawns []image.Point){
	x, y := GetSpawn(spawns)
	instance.X = x
	instance.Y = y
}

// func (p *PC) savePositionHistory() {
// 	p.PositionHistory.Add(struct{ X, Y float64 }{X: p.X, Y: p.Y})
// }

func (p *PC) SetX(x float64) {
	p.X = x
	// p.savePositionHistory()
}

func (p *PC) SetY(y float64) {
	p.Y = y
	// p.savePositionHistory()
}

// func (p *PC) IsXChanged() bool {
// 	var prevX float64
// 	for _, v := range p.PositionHistory.Get() {
// 		pos := v.(struct{ X, Y float64 })

// 		if prevX == 0 {
// 			prevX = pos.X
// 			continue
// 		}
// 		if prevX == pos.X {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (p *PC) IsYChanged() bool {
// 	var prevY float64
// 	for _, v := range p.PositionHistory.Get() {
// 		pos := v.(struct{ X, Y float64 })

// 		if prevY == 0 {
// 			prevY = pos.Y
// 			continue
// 		}
// 		if prevY == pos.Y {
// 			return false
// 		}
// 	}
// 	return true
// }

func (p *PC) SetSpeed(speedX float64) {
	p.Buffs.SpeedX = speedX
	wx, wy := screen.GetMaxWidth(), screen.GetMaxHeight()
	p.Buffs.SpeedY = speedX*float64(wy)/float64(wx) + .5
}

//Checks if pc model executes animation
func (p *PC) IsAnimatied() bool {
	return len(p.Physics.Jump) != 0
}

//Returns movement rotation related to the last
//movement direction
func (p *PC) GetMovementRotation() float64 {
	if history.GetDirection() == direction.LEFT {
		return -1
	}
	return 1
}

func UsePC() *PC {
	if instance == nil {
		instance = new(PC)

		id, err := uuid.NewUUID()
		if err != nil {
			logrus.Fatal("failed to create uuid for world:", err)
		}
		instance.ID = id

		instance.Username = EMPTY
		instance.PositionHistory = zeroshifter.New(2)

		instance.Health = DEFAULT_HEALTH
		instance.SetSpeed(1.5)
		instance.Equipment.Skin.Animation.FrameDelay = 5
		instance.Equipment.Skin.Animation.FrameDelayCounter = 1

		instance.Image = imagecollection.GetImage("assets/images/heroes/pumpkin")
		instance.Metadata = metadatacollection.GetMetadata("assets/images/heroes/pumpkin")
	}
	return instance
}
