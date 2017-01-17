package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App represents the application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize sets up the database connection and routes for the app
func (a *App) Initialize(user, password, dbname string) {}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {}
