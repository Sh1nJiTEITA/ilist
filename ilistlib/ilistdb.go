package ilistlib

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	ErrCantOpenBD           = errors.New("cant open DB")
	ErrCantCreateDB         = errors.New("cant create DB")
	ErrCantCreateDBIsDir    = errors.New("cant create DB input path points not to *.db file")
	ErrCantCreateUsersTable = errors.New("cant create users table")
	ErrCantPrepareUser      = errors.New("cant prepair insert user promt")
	ErrCantInsertUser       = errors.New("cant insert user")
)

func DBPath() string {
	abs, _ := filepath.Abs("./tasks.db")
	return abs
}

func createDB(path string) (*sql.DB, error) {
	if filepath.Ext(path) == "" {
		return nil, ErrCantCreateDBIsDir
	}
	if _, err := os.Stat(path); err != nil {
		file, err := os.Create(path)
		if err != nil {
			return nil, ErrCantCreateDB
		}
		file.Close()
		log.Infof("Database created at path %v", path)
	} else {
		log.Infof("Database exists at %v", path)
	}
	connection, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, ErrCantOpenBD
	}
	log.Info("Database connection is ok")
	return connection, nil
}

func CreateDB(path string) (*Database, error) {
	conn, err := createDB(path)
	if err != nil {
		return nil, err
	}
	db := Database{conn}
	return &db, nil
}

type User struct {
	Id       int64
	Username string
	password string
}

func (t User) String() string {
	return fmt.Sprintf("User(Id: %v, Username: %s)", t.Id, t.Username)
}

type Database struct {
	db *sql.DB
}

func (d *Database) AddUsersTable() error {
	promt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	_, err := d.db.Exec(promt)
	if err != nil {
		return ErrCantCreateUsersTable
	}
	log.Println("Users table was created")
	return nil
}

func (d *Database) AddUser(username string, password string) (*User, error) {
	promt := `INSERT INTO users (username, password) VALUES (?, ?)`
	stmt, err := d.db.Prepare(promt)
	if err != nil {
		return nil, ErrCantPrepareUser
	}
	defer stmt.Close()

	result, err := stmt.Exec(username, password)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	log.Printf("User '%v' with id '%v' added\n", username, id)
	return &User{id, username, password}, nil
}

func (d *Database) IsExistsUser(username string, password string) (bool, error) {
	promt := `SELECT COUNT(*) FROM users WHERE username = ? AND password = ?`
	stmt, err := d.db.Prepare(promt)
	if err != nil {
		return false, ErrCantPrepareUser
	}
	defer stmt.Close()
	var count int
	err = stmt.QueryRow(username, password).Scan(&count)
	if err != nil {
		return false, err
	}
	log.Printf("User '%v' with checked for existance\n (%v)", username, count > 0)
	return count > 0, nil
}

func (d *Database) DeleteUser(username string, password string) error {
	promt := `DELETE FROM users WHERE (username, password) VALUES (?, ?)`
	stmt, err := d.db.Prepare(promt)
	if err != nil {
		return ErrCantPrepareUser
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password)
	if err != nil {
		return err
	}
	log.Printf("User '%v' deleted\n", username)
	return nil
}

func (d *Database) GetUserByName(username string) (*User, error) {
	promt := `SELECT * FROM users WHERE username = ?`
	stmt, err := d.db.Prepare(promt)
	if err != nil {
		return nil, ErrCantPrepareUser
	}
	defer stmt.Close()

	var found User

	err = stmt.QueryRow(username).Scan(&found.Id, &found.Username, &found.password)
	if err != nil {
		return nil, err
	}
	log.Printf("User '%v' was found by name\n", found)
	return &found, nil
}
