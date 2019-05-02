package migrations

import (
	"log"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("adding uuid extension...")
		_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
		return err
	}, func(db migrations.DB) error {
		log.Println("dropping uuid extension...")
		_, err := db.Exec(`DROP EXTENSION "uuid-ossp"`)
		return err
	})
}
