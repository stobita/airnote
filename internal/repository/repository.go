package repository

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/olivere/elastic/v7"
	"github.com/volatiletech/sqlboiler/boil"
)

type repository struct {
	db         *sql.DB
	httpClient *http.Client
	esClient   *elastic.Client
}

// New return new repository
func New(db *sql.DB, httpClient *http.Client, esClient *elastic.Client) *repository {
	if os.Getenv("PRODUCTION") != "true" {
		boil.DebugMode = true
	}
	return &repository{
		db,
		httpClient,
		esClient,
	}
}
