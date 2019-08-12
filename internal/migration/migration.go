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
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	))
	if err != nil {
		return errors.Wrap(err, "Failed connection")
	}
	defer db.Close()

	return goose.Run("up", db, dir)
}
