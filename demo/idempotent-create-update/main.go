package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func ConnectDB(config Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s",
		config.User, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Order struct {
	ID      int64
	UserID  int64
	Status  string
	Version int
}

func CreateOrder(db *sql.DB, orderID int64, userID int64) (bool, error) {
	result, err := db.Exec("INSERT IGNORE INTO orders (id, user_id) VALUES (?, ?)", orderID, userID)
	if err != nil {
		return false, err
	}
	affectedRows, err := result.RowsAffected()
	return affectedRows != 0, err
}

func UpdateOrder(db *sql.DB, orderID int64, status string, version int) (bool, error) {
	result, err := db.Exec("UPDATE orders SET status = ?, version = version + 1 WHERE id = ? AND version = ?", status, orderID, version)
	if err != nil {
		return false, err
	}
	affectedRows, err := result.RowsAffected()
	return affectedRows != 0, err
}
