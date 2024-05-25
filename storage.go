package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Storage interface {
    GetRataRata() (float64, error)
}

type MysqlStore struct {
    db *sql.DB
}

func NewMysqlStore() (*MysqlStore, error) {
    dbUsername := os.Getenv("DBUSERNAME")
    dbPassword := os.Getenv("DBPASS")
    dbHost := os.Getenv("DBHOST")
    dbPort := os.Getenv("DBPORT")
    dbName := os.Getenv("DBNAME")

    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
      log.Fatal("Error connecting to database")
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &MysqlStore{
        db: db,
    }, nil
}

func (s *MysqlStore) GetRataRata() (float64, error) {
    query := "SELECT AVG(ipk) FROM ds_wisuda_tibil"
    var rata float64

    err := s.db.QueryRow(query).Scan(&rata)
    if err != nil {
        return 0, err
    }

    return rata, nil
}


