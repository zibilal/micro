package data_access

import (
	"github.com/mataharimall/micro-api/library/mysql"
	c "github.com/spf13/viper"
	"errors"
	"strings"
)

const (
	MYSQL = "mysqlaccess"
	MONGO = "mongoaccess"
)

func NewDataAccessFactory(dbapps string) (DataAccessor, error) {

	switch dbapps {
	case MYSQL:
		return NewMysqlRepository(mysql.Master(), mysql.Slave())
	case MONGO:
		database := c.GetString("mongo.database")
		server := c.GetString("mongo.master")
		return NewMongoRepository(server, database)
	}
	dbstrings := []string{
		MYSQL, MONGO,
	}

	return nil, errors.New("Invalid db apps " + dbapps + ", please use (" + strings.Join(dbstrings, "|")+ ")" )

}
