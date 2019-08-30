package repository

import (
	"database/sql"
	"net/http"
)

type repository struct {
	db         *sql.DB
	httpClient *http.Client
}

// New return new repository
func New(db *sql.DB, httpClient *http.Client) *repository {
	return &repository{
		db,
		httpClient,
	}
}
