package storage

import (
	"database/sql"
	"path/filepath"

	"github.com/YarikRevich/HideSeek-Client/internal/core/paths"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

type DB struct {
	db *sql.DB
}

func (d *DB) initTables() {

	if _, err := d.db.Exec("CREATE TABLE IF NOT EXISTS user (name VARCHAR(50) UNIQUE)"); err != nil {
		logrus.Fatal("initializing of user table failed: ", err)
	}

	if _, err := d.db.Exec("CREATE TABLE IF NOT EXISTS window (height REAL, width REAL)"); err != nil {
		logrus.Fatal("initializing of window table failed: ", err)
	}

}

func (d *DB) initDefaultValues() {
	q := d.db.QueryRow("SELECT COUNT(*) FROM user")
	var count int
	if err := q.Scan(&count); err != nil {
		logrus.Fatal("selecting count of rows user failed: ", err)
	}

	if _, err := d.db.Exec("INSERT INTO user (name) VALUES (?)", xid.New().String()); err != nil {
		logrus.Fatal("inserting default username failed:", err)
	}
}

func (d *DB) init() {
	d.initTables()
	d.initDefaultValues()
}

func NewDB() *sql.DB {
	d, err := sql.Open("sqlite3", filepath.Join(paths.GAME_STORAGE_DIR, "storage.db"))
	if err != nil {
		logrus.Fatal("connection to db failed: ", err)
	}

	session := &DB{d}
	session.init()
	
	return d
}
