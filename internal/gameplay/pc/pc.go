package pc

import (
	"image"

	"github.com/YarikRevich/HideSeek-Client/internal/direction"
	"github.com/YarikRevich/HideSeek-Client/internal/gameplay/objects"
	"github.com/YarikRevich/HideSeek-Client/internal/history"
	imagecollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/image_loader/collection"
	metadatacollection "github.com/YarikRevich/HideSeek-Client/internal/resource_manager/metadata_loader/collection"
	"github.com/YarikRevich/HideSeek-Client/internal/screen"
	"github.com/YarikRevich/HideSeek-Client/internal/storage/provider"
	"github.com/YarikRevich/caching/pkg/zeroshifter"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_HEALTH = 10
)

var instance *PC

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
	objects.Object
	
	Username string
	Health uint64
	//Spawn generated before the game start
	Spawn image.Point

	Team Team

	Buffs Buffs

	Physics   Physics
	Equipment Equipment
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
func (p *PC) SetSpawn(spawns []image.Point) {
	x, y := GetSpawn(spawns)
	instance.RawPos.X = x
	instance.RawPos.Y = y
}

func (p *PC) SetSpeed(speedX float64) {
	p.Buffs.SpeedX = speedX
	wx, wy := screen.GetMaxWidth(), screen.GetMaxHeight()
	p.Buffs.SpeedY = speedX*float64(wy)/float64(wx) * 2
}

//Returns position saved before animation processing
func (p *PC) GetPositionBeforeAnimation() (float64, float64) {
	return float64(p.Equipment.Skin.Animation.PositionBeforeAnimation.X),
		float64(p.Equipment.Skin.Animation.PositionBeforeAnimation.Y)
}

//Checks if pc model executes animation
func (p *PC) IsAnimatied() bool {
	return len(p.Physics.Jump) != 0
}

//Returns movement rotation related to the last
//movement direction
func (p *PC) GetMovementRotation() float64 {
	if history.GetDirection() == direction.LEFT && p.RawPos.X != 0 {
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

		instance.PositionHistory = zeroshifter.New(2)

		instance.Health = DEFAULT_HEALTH
		instance.SetSpeed(10)
		instance.Equipment.Skin.Animation.FrameDelay = 5
		instance.Equipment.Skin.Animation.FrameDelayCounter = 1

		instance.Image = imagecollection.GetImage("assets/images/heroes/pumpkin")
		instance.Metadata = metadatacollection.GetMetadata("assets/images/heroes/pumpkin")
	}
	return instance
}
