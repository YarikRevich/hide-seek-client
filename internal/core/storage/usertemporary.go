package storage

import "github.com/rs/xid"


type UserTemorary struct {
	name string
}

func (u *UserTemorary) GetUsername()string{
	if u.name == ""{
		u.name = xid.New().String()
	}
	return u.name
}

func (u *UserTemorary) SetUsername(name string){
	u.name = name
}

func NewUserTemorary()IUser{
	return new(UserTemorary)
}
