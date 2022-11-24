// Package postgres contains utilities for connecting to a postgres database server.
//
// Copyright © 2022 Karl Bateman. All Rights Reserved.
package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

const defaultDatabaseURI = "postgres://postgres:postgres@localhost"

// GetURI returns the database connection URI from the environment.
//
// If the environment variable `POSTGRES_DATABASE_URI` is not set, then a default is used.
func GetURI() string {
	uri := os.Getenv("DATABASE_URI")
	if uri == "" {
		uri = defaultDatabaseURI
	}
	return uri
}

// New is a factory for creating a Postgres database connection pool.
//
// See https://github.com/jackc/pgx/wiki/Getting-started-with-pgx#using-a-connection-pool for details.
func New() (*pgxpool.Pool, error) {
	uri := GetURI()
	return pgxpool.New(context.Background(), uri)
}
