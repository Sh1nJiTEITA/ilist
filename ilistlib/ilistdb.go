package ilistlib

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	// "github.com/Sh1nJiTEITA/ilist/utils"
	ut "github.com/Sh1nJiTEITA/ilist/utils"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	ErrCantOpenBD           = errors.New("cant open DB")
	ErrCantCreateDB         = errors.New("cant create DB")
	ErrCantCreateDBIsDir    = errors.New("cant create DB input path points not to *.db file")
	ErrCantCreateUsersTable = errors.New("cant create users table")
	ErrCantCreateTasksTable = errors.New("cant create tasks table")
	ErrCantPrepareUser      = errors.New("cant prepair insert user promt")
	ErrCantInsertUser       = errors.New("cant insert user")
	ErrCantUpdateUser       = errors.New("cant update user, user with input id does not exist")
)

var (
	IsCheckDBCreationPath = false
)

func DBPath() string {
	abs, _ := filepath.Abs("./tasks.db")
	return abs
}

func createDB(path string) (*sql.DB, error) {
	if IsCheckDBCreationPath {
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
	}
	connection, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, ErrCantOpenBD
	}
	log.Info("Database connection is ok")
	return connection, nil
}

func CreateDB(path string) (*Database, error) {
	conn := ut.Must(createDB(path))
	db := Database{conn}
	return &db, nil
}

type Task struct {
	Id      int64
	UserId  int64
	Content string
	Status  bool
	Level   int64
}

func (t Task) String() string {
	return fmt.Sprintf(`Task{Id: %v, UserId: %v, Status: %v, Content: %v}`, t.Id, t.UserId, t.Status, t.Content)
}

type Database struct {
	db *sql.DB
}

func (d *Database) AddUserTable() error {
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

func (d *Database) AddTaskTable() error {
	promt := `
	CREATE TABLE IF NOT EXISTS tasks (
		id	INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		status  INTEGER NOT NULL,
		level   INTEGER NOT NULL,
		
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err := d.db.Exec(promt)
	if err != nil {
		return ErrCantCreateTasksTable
	}
	log.Println("Tasks table was created")
	return nil
}

func (d *Database) AddTask(user *User, content string, status bool, level int64) (*Task, error) {
	promt := `
	INSERT INTO tasks (user_id, content, status, level) 
	VALUES (?, ?, ?, ?)
	`
	result, err := d.db.Exec(promt, user.Id, content, status, level)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}
	out_task := Task{
		Id:      id,
		UserId:  user.Id,
		Content: content,
		Status:  status,
		Level:   level,
	}

	log.Printf("%v added\n", out_task)
	return &out_task, nil
}

func (d *Database) IsExistsTask(id int64, unit_id int64) bool {
	promt := `SELECT * FROM tasks WHERE id = ? AND unit_id = ?`
	row := d.db.QueryRow(promt, id, unit_id)
	return row != nil
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

func (d *Database) UpdateUser(id int64, username string, password string) (*User, error) {
	promt := `UPDATE users SET username = ?, password = ? WHERE id = ?`
	stmt, err := d.db.Prepare(promt)
	if err != nil {
		return nil, ErrCantPrepareUser
	}
	defer stmt.Close()

	old_user, err := d.GetUserById(id)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(username, password, id)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, ErrCantUpdateUser
	}

	log.Printf("User updated: %v ===> %v", old_user, User{id, username, password})
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
	promt := `DELETE FROM users WHERE username = ? AND password = ?`
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

func (d *Database) GetUserById(id int64) (*User, error) {
	promt := `SELECT * FROM users WHERE id = ?`
	stmt, err := d.db.Prepare(promt)
	if err != nil {
		return nil, ErrCantPrepareUser
	}
	defer stmt.Close()

	var found User

	err = stmt.QueryRow(id).Scan(&found.Id, &found.Username, &found.password)
	if err != nil {
		return nil, err
	}
	log.Printf("User '%v' was found by id\n", found)
	return &found, nil
}

func (d *Database) GetTasksByUserId(user_id int64) ([]Task, error) {
	promt := `SELECT * FROM tasks WHERE user_id = ?`
	rows, err := d.db.Query(promt, user_id)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.Id, &t.UserId, &t.Content, &t.Status, &t.Level); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	log.Infof("Searching tasks by user_id... Have been found %v tasks", len(tasks))
	return tasks, nil
}

// === === === === === === === === === === === === === === === === === === ===
