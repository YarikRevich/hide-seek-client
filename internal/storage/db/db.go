package db

import (
	"database/sql"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/YarikRevich/HideSeek-Client/internal/paths"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

func NewDB() *sql.DB {
	c, err := sql.Open("sqlite3", filepath.Join(paths.GAME_STORAGE_DIR, "storage.db"))
	if err != nil {
		logrus.Fatal("connection to db failed: ", err)
	}
	go func(){
		n := make(chan os.Signal, 1)
		signal.Notify(n, os.Interrupt)
		for range n{
			if err := c.Close(); err != nil{
				logrus.Fatal("closing of db failed: ", err)
			}
		}
	}()

	if _, err := c.Exec("CREATE TABLE IF NOT EXISTS user (name VARCHAR(50) UNIQUE)"); err != nil{
		logrus.Fatal("initializing of user table failed: ", err)
	}

	if _, err = c.Exec("CREATE TABLE IF NOT EXISTS window (height REAL, width REAL)"); err != nil{
		logrus.Fatal("initializing of window table failed: ", err)
	}

	q := c.QueryRow("SELECT COUNT(*) FROM user")
	var count int
	if err := q.Scan(&count); err != nil{
		logrus.Fatal("selecting count of rows user failed: ", err)
	}

	if _, err = c.Exec("INSERT INTO user (name) VALUES (?)", xid.New().String()); err != nil{
		logrus.Fatal("inserting default username failed:", err)
	}

	return c
}
