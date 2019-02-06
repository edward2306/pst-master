package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/pisatoo/pst-master/database"
	"github.com/spf13/viper"
)

func loadConfig() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()

	return err
}

func runServer(db *sql.DB) {
	r := LoadRouter(db)

	log.Println("Server run on " + getAddress())
	http.ListenAndServe(getAddress(), r)
}

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.InitDatabase()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	runServer(db)

	defer db.Close()
}
