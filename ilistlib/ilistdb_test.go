package ilistlib

import (
	"testing"
)

func checkErr(err error, t *testing.T) bool {
	if err != nil {
		t.Fatal(err)
		return false
	}
	return true
}

func Test_CreateDB(t *testing.T) {
	_, err := CreateDB(":memory:")
	checkErr(err, t)
}

func Test_AddUserTable(t *testing.T) {
	db, err := CreateDB(":memory:")
	checkErr(err, t)
	err = db.AddUsersTable()
	checkErr(err, t)
}

func Test_AddUser(t *testing.T) {
	db, err := CreateDB(":memory:")
	checkErr(err, t)

	err = db.AddUsersTable()
	checkErr(err, t)

	user, err := db.AddUser("some", "123")
	checkErr(err, t)
	if user.Id != 1 || user.Username != "some" || user.password != "123" {
		t.Fatalf(`Invalid one of next params (check 1):
			 %v !=? 1,
			 %v !=? "some"
			 %v !=? "123"
			 `, user.Id, user.Username, user.password)
	}
	user, err = db.GetUserByName("some")
	checkErr(err, t)
	if user.Id != 1 || user.Username != "some" || user.password != "123" {
		t.Fatalf(`Invalid one of next params (check 2):
			 %v !=? 1,
			 %v !=? "some"
			 %v !=? "123"
			 `, user.Id, user.Username, user.password)
	}
}

func Test_DeleteUser(t *testing.T) {
	db, err := CreateDB(":memory:")
	checkErr(err, t)

	err = db.AddUsersTable()
	checkErr(err, t)

	user, err := db.AddUser("some", "123")
	checkErr(err, t)
	if user.Id != 1 || user.Username != "some" || user.password != "123" {
		t.Fatalf(`Invalid one of next params (check 1):
			 %v !=? 1,
			 %v !=? "some"
			 %v !=? "123"
			 `, user.Id, user.Username, user.password)
	}
	err = db.DeleteUser("some", "123")
	checkErr(err, t)
	if status, err := db.IsExistsUser("some", "123"); checkErr(err, t) && status {
		t.Fatalf("User wasn't deleted...")
	}
}
