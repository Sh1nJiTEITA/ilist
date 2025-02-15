package ilistlib

import "testing"

func checkErr(err error, t *testing.T) bool {
	if err != nil {
		t.Fatal(err)
		return false
	}
	return true
}

func TestOpenDB(t *testing.T) {
	_, err := OpenDB(":memory:")
	checkErr(err, t)
}

func TestAddUserTable(t *testing.T) {
	db, err := OpenDB(":memory:")
	checkErr(err, t)

	var user_table TableUsers
	err = db.CreateTable(user_table)
	checkErr(err, t)

	if !db.IsTableExists(user_table) {
		t.Fatal("User table does not exist")
	}
	if err = db.CreateTable(user_table); err != ErrDbTableExists {
		t.Fatal("User table must exist already but not")
	}
}

func TestAddUser(t *testing.T) {
	db, _ := OpenDB(":memory:")
	user_table := NewTableUsers(db)
	db.CreateTable(user_table)

	user := NewUser("snj", "1488")

	err := user_table.Save(user)
	if err != nil {
		t.Fatal("cant save user")
	}
	found_user, err := user_table.FindByName("snj")
	if err != nil {
		t.Fatal(err)
	}
	if found_user.Id != 1 {
		t.Fatal("New user id is not 1")
	}
	if user.Username != found_user.Username || user.Password != found_user.Password {
		t.Fatal("Current and added user are not the same")
	}

	all, err := user_table.GetAll()
	checkErr(err, t)
	if len(all) != 1 {
		t.Fatal("Too many users")
	}
	if all[0].Id != 1 || all[0].Username != "snj" || all[0].Password != "1488" {
		t.Fatal("Got user via get all are incorrect")
	}

	user2 := NewUser("snj2", "228")

	err = user_table.Save(user2)
	checkErr(err, t)

	found_user, err = user_table.FindByName("snj2")
	if err != nil {
		t.Fatal(err)
	}
	if found_user.Id != 2 {
		t.Fatal("New user id is not 1")
	}
	if user2.Username != found_user.Username || user2.Password != found_user.Password {
		t.Fatal("Current and added user are not the same")
	}

	all, err = user_table.GetAll()
	checkErr(err, t)
	if len(all) != 2 {
		t.Fatal("Too many users")
	}
	if all[0].Id != 1 || all[0].Username != "snj" || all[0].Password != "1488" {
		t.Fatal("Got user via get all are incorrect")
	}

	if all[1].Id != 2 || all[1].Username != "snj2" || all[1].Password != "228" {
		t.Fatal("Got user via get all are incorrect")
	}
	if all[1].Id != 2 || all[1].Username != user2.Username || all[1].Password != user2.Password {
		t.Fatal("Got user via get all are incorrect")
	}

}
