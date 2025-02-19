package ilistlib

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	ErrDbTableExists = errors.New("table exists")
)

type Table interface {
	TableCreatePromt() string
	TableName() string
}

type IDB struct {
	db *sql.DB
}

func OpenDB(path string) (*IDB, error) {
	connection, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("Cant open db due to driver error: %w", err)
	}
	log.Info("Database connection is ok")
	return &IDB{db: connection}, nil
}

func (db *IDB) IsTableExists(table Table) bool {
	promt := "PRAGMA table_info(%s);"
	rows, err := db.db.Query(fmt.Sprintf(promt, table.TableName()))
	if err != nil {
		return false
	}
	defer rows.Close()
	return rows.Next()
}

func (db *IDB) CreateTable(table Table) error {
	if db.IsTableExists(table) {
		return ErrDbTableExists
	}
	_, err := db.db.Exec(table.TableCreatePromt())
	if err != nil {
		return err
	}
	return nil
}
