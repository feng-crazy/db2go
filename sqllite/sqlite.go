package sqlite

import (
	"database/sql"
	"fmt"

	_ "gorm.io/driver/sqlite"
)

type Sqlite struct {
	DB     *sql.DB
	DBName string
}

func (this *Sqlite) Init(userName, pwd, host string, port int, dbName string) error {
	_db, err := sql.Open("mysql", dbName)
	if err != nil {
		return err
	}
	this.DB = _db
	this.DBName = dbName
	return nil
}
func (this *Sqlite) FindTableString() string {
	return "SELECT name FROM sqlite_master where type='table' order by name"
}
func (*Sqlite) FindColumnsString() string {
	return findColumnSql
}

func (this *Sqlite) DBNameString() string {
	return this.DBName
}

func (this *Sqlite) GetDB() *sql.DB {
	return this.DB
}
