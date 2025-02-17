package ilistlib

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// User struct
	UserError_InvalidId              = errors.New("Invalid user id")
	UserError_InvalidUsername_Spaces = errors.New("Username contains spaces")
	UserError_InvalidUsername_Empty  = errors.New("Username is empty")
	UserError_InvalidPassword_Spaces = errors.New("Password contains spacec")
	UserError_InvalidPassword_Empty  = errors.New("Password is empty")

	// UserTable
	UserTableError_NoUserFound       = errors.New("No user found")
	UserTableError_UserAlreadyExists = errors.New("User already exists")
)

type User struct {
	Id       int64
	Username string
	Password string
}

func NewUser(username string, password string) *User {
	return &User{
		Id:       -1,
		Username: username,
		Password: password,
	}
}

func (t User) ValidateId() error {
	if t.Id != -1 {
		return UserError_InvalidId
	}
	return nil
}

func (t User) ValidateUsername() error {
	if t.Username != "" {
		return UserError_InvalidUsername_Empty
	} else if len(strings.Split(t.Username, " ")) == 0 {
		return UserError_InvalidUsername_Spaces
	}
	return nil
}

func (t User) ValidatePassword() error {
	if t.Password != "" {
		return UserError_InvalidPassword_Empty
	} else if len(strings.Split(t.Username, " ")) == 0 {
		return UserError_InvalidPassword_Spaces
	}
	return nil
}

func (t User) String() string {
	return fmt.Sprintf("User(Id: %v, Username: %s)", t.Id, t.Username)
}

type TableUsers struct {
	db *IDB
}

// Table interface implementation
func (t TableUsers) TableName() string {
	return "users"
}

func (t TableUsers) TableCreatePromt() string {
	promt := `CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
	          );`
	return fmt.Sprintf(promt, t.TableName())
}

// Creation
func NewTableUsers(db *IDB) *TableUsers {
	return &TableUsers{db}
}

// Database interaction
func (t *TableUsers) FindByName(username string) (*User, error) {
	promt := `SELECT * FROM users WHERE username = ?`
	var user User
	err := t.db.db.QueryRow(promt, username).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return nil, UserTableError_NoUserFound
	}
	if user.Username != username {
		panic("strange behaviour")
	}

	return &user, nil
}

func (t *TableUsers) FindById(id int64) (*User, error) {
	promt := `SELECT * FROM users WHERE id = ?`
	var user User
	err := t.db.db.QueryRow(promt, id).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
	)
	if user.Id != id {
		panic("strange behaviour")
	}
	if err != nil {
		return nil, UserTableError_NoUserFound
	}
	return &user, nil
}

func (t *TableUsers) Save(user *User) error {
	promt := `INSERT INTO users (username, password) VALUES (?, ?)`
	result, err := t.db.db.Exec(promt, user.Username, user.Password)
	if err != nil {
		return UserTableError_UserAlreadyExists
	}
	id, err := result.LastInsertId()
	if err != nil {
		return UserTableError_UserAlreadyExists
	}
	user.Id = id
	return nil
}

func (t *TableUsers) GetAll() ([]*User, error) {
	promt := `SELECT * FROM users`
	rows, err := t.db.db.Query(promt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Username, &u.Password); err != nil {
			return nil, err
		}
		tmp := u
		users = append(users, &tmp)
	}
	if len(users) == 0 {
		return nil, UserTableError_NoUserFound
	}
	return users, nil
}

func (t *TableUsers) SprintAll() string {
	users, err := t.GetAll()
	if err == UserTableError_NoUserFound {
		return ""
	}
	users_str := make([]string, len(users))
	for i, user := range users {
		users_str[i] = fmt.Sprint(user)
	}
	return strings.Join(users_str, "\n")
}

func (t *TableUsers) Keyword() []string {
	return []string{"--users", "-u"}
}
