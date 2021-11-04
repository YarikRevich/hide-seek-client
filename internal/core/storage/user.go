package storage

import (
	"database/sql"
	"strings"
	"github.com/sirupsen/logrus"
)

type User struct {
	db *sql.DB
}

func (u *User) GetUsername()string{
	q, err := u.db.Query("SELECT name FROM user")
	if err != nil {
		logrus.Fatal("selecting username failed: ", err)
	}

	var name string
	if q.Next() {
		if err := q.Scan(&name); err != nil {
			logrus.Fatal("scaning username failed: ", err)
		}
	}
	if err = q.Close(); err != nil {
		logrus.Fatal("failed closing query: ", err)
	}

	return name
}

func (u *User) SetUsername(name string){
	q := NewQuery()
	q.AddField("name")
	q.AddValue(name)

	r := strings.Repeat("?,", len(q.Values))
	if _, err := u.db.Exec("DELETE FROM user"); err != nil {
		logrus.Fatal(err)
	}
	if _, err := u.db.Exec("INSERT INTO user("+q.GetFieldsAsString()+") VALUES ("+r[:len(r)-1]+")", q.Values...); err != nil {
		logrus.Fatal(err)
	}
}

func NewUser(db *sql.DB)*User{
	return &User{db}
}
