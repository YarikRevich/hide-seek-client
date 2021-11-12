package storage

import "github.com/YarikRevich/HideSeek-Client/tools/cli"

var instance *Storage


type Storage struct {
	user IUser
}

func (s *Storage) User() IUser{
	return s.user
}

func UseStorage() *Storage{
	if instance == nil{
		db := NewDB()
		instance = new(Storage)
		if cli.IsDisableConfigAutoSave(){
			instance.user = NewUserTemorary()
		}else{
			instance.user = NewUserStorage(db)
		}
	}
	return instance
}