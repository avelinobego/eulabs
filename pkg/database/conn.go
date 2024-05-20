package database

import (
	config "eulabs/products/etc"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDatabase(cfg *config.Config) (db *sqlx.DB, err error) {
	// Create the connection string.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)

	// Connect to the MySQL database.
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the database connection.
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db.Unsafe(), nil
}
