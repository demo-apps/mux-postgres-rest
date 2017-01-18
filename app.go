package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App represents the application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize sets up the database connection and routes for the app
func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {}
