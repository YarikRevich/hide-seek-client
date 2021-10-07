package user

import (
	"database/sql"
	"strings"

	"github.com/YarikRevich/HideSeek-Client/internal/storage/common"
	"github.com/sirupsen/logrus"
)

type user struct {
	db *sql.DB
}

func (u *user) Get(f string) interface{} {
	return nil
}

func (u *user) Save(q common.DBQuery) {
	v := q.FormattedValues()
	r := strings.Repeat("?,", len(v))
	if _, err := u.db.Exec("DELETE FROM user"); err != nil{
		logrus.Fatal(err)
	}
	if _, err := u.db.Exec("INSERT INTO user("+q.FormattedFields()+") VALUES (" + r[:len(r)-1] + ")", v...); err != nil {
		logrus.Fatal(err)
	}
}

func NewUserStorage(db *sql.DB) common.StorageBlock {
	return &user{db}
}
