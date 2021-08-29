package user

const (
	EMPTY = ""
	DEFAULT_HEALTH = 10
)

var (
	user *User
)

type Animation struct {
	Updation uint32
	UpdationDelay uint32
	CurrentFrame uint32
	CurrentFrameMatrix []float64
}

type GameCredentials struct {
	LobbyID string
}

type Skin struct {
	Name string
	Animation Animation
}

type Weapon struct {
	Name string
	Radius int
	Animation Animation
}

type Equipment struct {
	Skin Skin
	Weapon string
}

type User struct{
	Username string

	Health uint64

	X float64
	Y float64

	Equipment Equipment

	GameCredentials GameCredentials
}

// userConfig := Users.User{
// 	// Conn: conn,
// 	// Pos: &Users.Pos{
// 		X: int(randomSpawn.X),
// 		Y: int(randomSpawn.Y),
// 	// },
// 	// GameInfo: &Users.GameInfo{
// 	// 	Health: 10,
// 	// 	WeaponName:  Utils.GetRandomWeaponImage(winConf.Components.AvailableWeaponImages),
// 	// },
// 	// PersonalInfo: &Users.PersonalInfo{
// 	// 	Username:    username,
// 	// 	HeroPicture: Utils.GetRandomHeroImage(winConf.Components.AvailableHeroImages),
// 	// },
// 	Animation:  &Users.Animation{CurrentFrameMatrix: []float64{0, 0, 0, 0}},
// 	// Networking: new(Users.Networking),
// 	// Context:    new(Users.Context),
// }


// type Networking struct{
// 	Index int
// }

func (u *User) SetUsername(un string){
	u.Username = un
}

func GetInstance()*User{
	if user == nil{
		return &User{
			Username: EMPTY,
			Health: DEFAULT_HEALTH,
		}
	}
	return user
}


