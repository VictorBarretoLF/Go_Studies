package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Connection string for MySQL
	dsn := "root:root@tcp(mysql:3306)/courses"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	defer db.Close()

	// Verify the connection to the database
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could not create database driver: %v\n", err)
	}

	// Import the file driver
	sourceURL := "file:///go/src/sql/migrations"
	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		"mysql", driver)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v\n", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while running the migrations: %v\n", err)
	}

	log.Println("Migrations applied successfully!")
}
