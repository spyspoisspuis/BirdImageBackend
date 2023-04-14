package main

import (
	"bird_golang_back/internal/db"
	"bird_golang_back/internal/router"
	"bird_golang_back/internal/util"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	util.InitViper()
	db.ConnectDB()
	defer db.DisconnectDB()


	r := router.RouterEngine()

	r.Run(fmt.Sprintf(":%d", viper.GetInt("connection.appPort")))
}
