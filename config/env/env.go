package env

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/joho/godotenv"
)

func Get() map[string]string {

	env, err := godotenv.Read()

	if err != nil {
		log.Info(fmt.Sprintf("error cause: %v", err.Error()))
		panic(err.Error())
	}

	return env
}
