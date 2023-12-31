package module

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() *gorm.DB {

	dsn := getDbDsn()
	
	var err error
	PostgresDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to postgres: ", err)
		os.Exit(1)
	}

	return PostgresDb
}

func getDbDsn() string {

	return "host=" + config["postgres.host"]+
		" port=" + config["postgres.port"] +
		" user=" + config["postgres.user"] +
		" password=" + config["postgres.password"] +
		" dbname=" + config["postgres.dbname"] +
		" TimeZone=Europe/Rome" +
		" sslmode=disable"

}
