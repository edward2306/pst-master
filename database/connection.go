package database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDatabase() (db *gorm.DB, err error) {
	dbDriver := viper.GetString("DB_DRIVER")
	var connectionString string

	if dbDriver == "mysql" {
		connectionString = buildMySqlConnectionString()
	} else if dbDriver == "postgres" {
		connectionString = buildPostgreSqlConnectionString()
	} else if dbDriver == "sqlite3" {
		connectionString = buildSqliteConnectionString()
	}

	db, err = openConnection(dbDriver, connectionString)

	return
}

func openConnection(dbDriver, connection string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbDriver, connection)

	if err != nil {
		fmt.Println(err)
	}

	return
}

func buildMySqlConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"))

	return
}

func buildPostgreSqlConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_POSTGRES_SSL_MODE"))

	return
}

func buildSqliteConnectionString() (connectionString string) {
	connectionString = viper.GetString("DB_NAME")
	return
}
