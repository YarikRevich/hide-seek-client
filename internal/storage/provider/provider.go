package provider

import (
	"github.com/YarikRevich/HideSeek-Client/internal/storage/common"
	"github.com/YarikRevich/HideSeek-Client/internal/storage/db"
	"github.com/YarikRevich/HideSeek-Client/internal/storage/user"
)

var instance IProvider

type IProvider interface {
	User() common.StorageBlock
}

type provider struct {
	userStorage common.StorageBlock
}

func (pr *provider) User() common.StorageBlock {
	return pr.userStorage
}

func UseStorageProvider() IProvider {
	if instance == nil {
		d := db.NewDB()

		instance = &provider{
			userStorage: user.NewUserStorage(d),
		}
	}
	return instance
}
