package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func logfatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB is...
func ConnectDB() *gorm.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	fmt.Println(pgURL)
	logfatal(err)

	db, err := gorm.Open(postgres.Open(pgURL), &gorm.Config{})
	logfatal(err)

	return db
}
