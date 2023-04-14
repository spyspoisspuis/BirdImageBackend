package db

import (
	"context"
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var database *sql.DB

func ConnectDB() {

	address := viper.GetString("connection.dbURL")

	db, err := sql.Open("mysql", address)
	if err != nil {
		panic("Cannot access mariadb server")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	database = db

}
func DisconnectDB() {
	database.Close()
}

func GetDatabase() *sql.DB {
	return database
}
