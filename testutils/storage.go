package testutils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/pressly/goose"
)

var testDB *sql.DB

func StorageTestRunner(m *testing.M) (int, error) {
	dbClose := setupTestDBConn()
	defer dbClose()
	// migrateTestDB()

	result := m.Run()

	truncateTables()
	return result, nil
}

func setupTestDBConn() func() {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	))
	if err != nil {
		log.Fatalf("Failed connect to mysql: %s", err)
	}
	testDB = db
	return func() { testDB.Close() }
}

func migrateTestDB() {
	err := goose.Run("up", testDB, "../../db/migrations")
	if err != nil {
		log.Fatalf("Failed migration: %s", err)
	}
}

func GetTestDBConn() (*sql.DB, func()) {
	if testDB == nil {
		log.Fatal("sql connection not initialized")
	}
	return testDB, func() { truncateTables() }
}

func truncateTables() {
	rows, err := testDB.Query("SHOW TABLES;")
	if err != nil {
		log.Fatalf("show tables error: %s", err)
	}
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("show tables error: %s", err)
		}
		cmds := []string{
			"SET FOREIGN_KEY_CHECKS = 0",
			fmt.Sprintf("TRUNCATE %s", tableName),
			"SET FOREIGN_KEY_CHECKS = 1",
		}
		for _, cmd := range cmds {
			if _, err := testDB.Exec(cmd); err != nil {
				log.Fatalf("show tables error: %s", err)
				continue
			}
		}
	}

}
