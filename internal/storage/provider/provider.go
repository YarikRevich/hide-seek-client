package provider

import (
	"github.com/YarikRevich/HideSeek-Client/internal/storage/common"
	"github.com/YarikRevich/HideSeek-Client/internal/storage/db"
	"github.com/YarikRevich/HideSeek-Client/internal/storage/user"
)

var instance IProvider

type IProvider interface {
	GetUsername() string
	SetUsername(username string) 
}

type provider struct {
	userStorage common.StorageBlock
}

func (pr *provider) GetUsername() string {
	return pr.userStorage.Get("name").(string)
}

func (pr *provider) SetUsername(username string)  {
	pr.userStorage.Save(common.DBQuery{{
		Field: "name", Value: username,
	}})
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
