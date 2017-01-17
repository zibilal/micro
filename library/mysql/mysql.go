package mysql

import (
	_ "database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	c "github.com/spf13/viper"
)

var mysqlConfig = map[string]*DB{
	"master": nil,
	"slave":  nil,
}

type DB struct {
	db *sqlx.DB
}

var lock = sync.RWMutex{}

func Master() (db *DB) {
	if mysqlConfig["master"] != nil {
		return mysqlConfig["master"]
	}
	lock.Lock()
	db = createMysqlConn("master")
	defer lock.Unlock()

	return
}

func Slave() (db *DB) {
	if mysqlConfig["slave"] != nil {
		return mysqlConfig["slave"]
	}
	lock.Lock()
	db = createMysqlConn("slave")
	defer lock.Unlock()

	return
}

func createMysqlConn(types string) *DB {
	db, err := sqlx.Open("mysql", setParams(c.GetStringMapString("mysql."+types)))
	if err != nil {
		log.Println("DB Connection Error: ", err)
		os.Exit(1)
	}
	//set max idle conn to preserve connection pool
	db.SetMaxIdleConns(c.GetInt("mysql.maxIdleConns"))
	//limit connection pool
	db.SetMaxOpenConns(c.GetInt("mysql.maxOpenConns"))
	return &DB{db}
}

func setParams(param map[string]string) (res string) {
	var h, l, p string
	d := c.GetString("mysql.database")
	if param == nil {
		return
	}
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", l, p, h, d)
}
