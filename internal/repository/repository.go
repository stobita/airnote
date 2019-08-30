package repository

import (
	"database/sql"
	"net/http"

	"github.com/go-redis/redis"
)

type repository struct {
	db          *sql.DB
	httpClient  *http.Client
	redisClient *redis.Client
}

// New return new repository
func New(db *sql.DB, httpClient *http.Client, redisClient *redis.Client) *repository {
	return &repository{
		db,
		httpClient,
		redisClient,
	}
}
