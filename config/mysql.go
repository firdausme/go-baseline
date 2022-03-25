package config

import (
	"database/sql"
	"fmt"
	log "github.com/jeanphorn/log4go"
	Env "go-baseline/config/env"
	c "go-baseline/constant"
	"go-baseline/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() *sql.DB {

	env := Env.Get()

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		env[c.DbUsername],
		env[c.DbPassword],
		env[c.DbHost],
		env[c.DbPort],
		env[c.DbName],
	)

	db, err := sql.Open(env[c.DbDriver], dbSource)
	if err != nil {
		_ = log.Error("Error validating sql.Open arguments ..")
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		_ = log.Error("Error verifying connection with db.Ping ..")
		panic(err.Error())
	}

	db.SetConnMaxIdleTime(time.Minute * time.Duration(utils.StrToInt(env[c.DbConnMaxIdleTime])))
	db.SetMaxOpenConns(utils.StrToInt(env[c.DbMaxOpenConns]))
	db.SetMaxIdleConns(utils.StrToInt(env[c.DbMaxIdleConns]))

	return db
}
