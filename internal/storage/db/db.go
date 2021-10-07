package db

import (
	"database/sql"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func NewDB() *sql.DB {
	c, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		logrus.Errorf("connection to db failed: %w", err)
	}
	go func(){
		n := make(chan os.Signal)
		signal.Notify(n, os.Interrupt)
		for range n{
			if err := c.Close(); err != nil{
				logrus.Errorf("closing of db failed: %w", err)
			}
		}
	}()

	if _, err := c.Exec("CREATE TABLE IF NOT EXISTS user (name VARCHAR(50) UNIQUE)"); err != nil{
		logrus.Errorf("initializing of user table failed: %w", err)
	}

	if _, err := c.Exec("CREATE TABLE IF NOT EXISTS window (height REAL, width REAL)"); err != nil{
		logrus.Errorf("initializing of window table failed: %w", err)
	}

	return c
}
