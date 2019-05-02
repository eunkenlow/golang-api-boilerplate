package migrations

import (
	"flag"
	"fmt"
	"log"

	"github.com/eunkenlow/golang-api-boilerplate/pkg/postgresql"
	"github.com/go-pg/migrations"
)

func init() {
	flag.Parse()

	db := postgresql.GetDb()

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatalf(err.Error(), err)
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}
