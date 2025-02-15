package ilistlib

//
// import (
// 	"testing"
// )
//
// func Test_CreateDB(t *testing.T) {
// 	_, err := CreateDB(":memory:")
// 	checkErr(err, t)
// }
//
// func Test_AddUserTable(t *testing.T) {
// 	db, err := CreateDB(":memory:")
// 	checkErr(err, t)
// 	err = db.AddUserTable()
// 	checkErr(err, t)
// }
//
// func Test_AddUser(t *testing.T) {
// 	db, err := CreateDB(":memory:")
// 	checkErr(err, t)
//
// 	err = db.AddUserTable()
// 	checkErr(err, t)
//
// 	user, err := db.AddUser("some", "123")
// 	checkErr(err, t)
// 	if user.Id != 1 || user.Username != "some" || user.password != "123" {
// 		t.Fatalf(`Invalid one of next params (check 1):
// 			 %v !=? 1,
// 			 %v !=? "some"
// 			 %v !=? "123"
// 			 `, user.Id, user.Username, user.password)
// 	}
// 	user, err = db.GetUserByName("some")
// 	checkErr(err, t)
// 	if user.Id != 1 || user.Username != "some" || user.password != "123" {
// 		t.Fatalf(`Invalid one of next params (check 2):
// 			 %v !=? 1,
// 			 %v !=? "some"
// 			 %v !=? "123"
// 			 `, user.Id, user.Username, user.password)
// 	}
// }
//
// func Test_DeleteUser(t *testing.T) {
// 	db, err := CreateDB(":memory:")
// 	checkErr(err, t)
//
// 	err = db.AddUserTable()
// 	checkErr(err, t)
//
// 	user, err := db.AddUser("some", "123")
// 	checkErr(err, t)
// 	if user.Id != 1 || user.Username != "some" || user.password != "123" {
// 		t.Fatalf(`Invalid one of next params (check 1):
// 			 %v !=? 1,
// 			 %v !=? "some"
// 			 %v !=? "123"
// 			 `, user.Id, user.Username, user.password)
// 	}
// 	err = db.DeleteUser("some", "123")
// 	checkErr(err, t)
// 	if status, err := db.IsExistsUser("some", "123"); checkErr(err, t) && status {
// 		t.Fatalf("User wasn't deleted...")
// 	}
// }
//
// func Test_FindUserById(t *testing.T) {
// 	db, err := CreateDB(":memory:")
// 	checkErr(err, t)
// 	err = db.AddUserTable()
// 	checkErr(err, t)
// 	user, err := db.AddUser("some", "123")
// 	checkErr(err, t)
//
// 	user_, err := db.GetUserById(user.Id)
// 	checkErr(err, t)
//
// 	if user.Id != user_.Id || user.Username != user_.Username || user.password != user_.password {
// 		t.Fatalf(`Invalid one of next params (check 1):
// 			 %v !=? 1,
// 			 %v !=? "some"
// 			 %v !=? "123"
// 			 `, user.Id, user.Username, user.password)
// 	}
// }
//
// func Test_FindUserByName(t *testing.T) {
// 	db, err := CreateDB(":memory:")
// 	checkErr(err, t)
// 	err = db.AddUserTable()
// 	checkErr(err, t)
// 	user, err := db.AddUser("some", "123")
// 	checkErr(err, t)
//
// 	user_, err := db.GetUserByName(user.Username)
// 	checkErr(err, t)
//
// 	if user.Id != user_.Id || user.Username != user_.Username || user.password != user_.password {
// 		t.Fatalf(`Invalid one of next params (check 1):
// 			 %v !=? 1,
// 			 %v !=? "some"
// 			 %v !=? "123"
// 			 `, user.Id, user.Username, user.password)
// 	}
// }
//
// func Test_AddTaskTable(t *testing.T) {
// 	db, err := CreateDB(":memory:")
// 	checkErr(err, t)
// 	err = db.AddTaskTable()
// 	checkErr(err, t)
// }
//
// func Test_AddTask(t *testing.T) {
// 	db, err := CreateDB(":memory:")
// 	checkErr(err, t)
// 	err = db.AddTaskTable()
// 	err = db.AddUserTable()
// 	checkErr(err, t)
// 	user, err := db.AddUser("some", "123")
// 	checkErr(err, t)
// 	task, err := db.AddTask(user, "SOME_CONTENT", false, 0)
// 	checkErr(err, t)
//
// 	if task.Id != 1 || task.UserId != user.Id || task.Content != "SOME_CONTENT" || task.Status != false || task.Level != 0 {
// 		t.Fatalf(`Invalid one of next params (check 1):
// 			 %v !=? 0,
// 			 %v !=? %v,
// 			 %v !=? "SOME_CONTENT"
// 			 %v !=? false
// 			 %v !=? 0
// 			 `, task.Id, task.UserId, user.Id, task.Content, task.Status, task.Level)
// 	}
// }
//
// func Test_GetTasks(t *testing.T) {
// 	db, err := CreateDB(":memory:")
// 	checkErr(err, t)
// 	err = db.AddTaskTable()
// 	err = db.AddUserTable()
// 	checkErr(err, t)
// 	user, err := db.AddUser("some", "123")
// 	checkErr(err, t)
// 	task_1, err := db.AddTask(user, "SOME_CONTENT_1", false, 0)
// 	checkErr(err, t)
// 	task_2, err := db.AddTask(user, "SOME_CONTENT_2", true, 0)
// 	checkErr(err, t)
// 	tasks, err := db.GetTasksByUserId(user.Id)
// 	checkErr(err, t)
//
// 	if tasks[0] != *task_1 || tasks[1] != *task_2 {
// 		t.Fatalf(`Tasks are different!
// 			 (needed : got):
// 			 %v : %v,
// 			 %v : %v
// 			`, task_1, tasks[0], task_2, tasks[1])
// 	}
// }
