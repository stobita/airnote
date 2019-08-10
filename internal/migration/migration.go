package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

const dir = "./db/migrations"

func Run() error {
	driver := "mysql"
	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open(driver, fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		return errors.Wrap(err, "Failed connection")
	}
	defer db.Close()

	return goose.Run("up", db, dir)
}
