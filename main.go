package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/pisatoo/pst-master/database"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
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
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	cliApp := cli.NewApp()
	cliApp.Name = "PISATOO"
	cliApp.Version = "1.0.0"

	cliApp.Commands = []cli.Command{
		{
			Name:  "migrate",
			Usage: "Run database migration",
			Action: func(c *cli.Context) error {
				err = database.Migrate(db)
				return err
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// runServer(db)

	defer db.Close()
}
