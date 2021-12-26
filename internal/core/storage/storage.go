package storage

import "github.com/YarikRevich/hide-seek-client/tools/params"

var instance *Storage

type Storage struct {
	user IUser
}

func (s *Storage) User() IUser {
	return s.user
}

func UseStorage() *Storage {
	if instance == nil {
		db := NewDB()
		instance = new(Storage)
		if params.IsDisableConfigAutoSave() {
			instance.user = NewUserTemorary()
		} else {
			instance.user = NewUserStorage(db)
		}
	}
	return instance
}
