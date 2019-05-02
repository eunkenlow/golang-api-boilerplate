package migrations

import (
	"log"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("creating table users...")
		_, err := db.Exec(`CREATE TABLE users(
			id varchar(48) PRIMARY KEY,
			email varchar(255) UNIQUE NOT NULL,
			name varchar(255),
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		)`)
		return err
	}, func(db migrations.DB) error {
		log.Println("dropping table users...")
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}
