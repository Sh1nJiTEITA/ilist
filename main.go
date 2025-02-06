package main

import (
	// "fmt"
	"github.com/Sh1nJiTEITA/ilist/ilistlib"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"os"
)

func SetupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// task := ilistlib.Task{false, "some task"}

	// tasks := make([]ilistlib.Task, 2)
	// tasks = append(tasks, ilistlib.Task{})
	// tasks = append(tasks, ilistlib.Task{})
	// ilistlib.ReadTasksFromFile(&tasks, "./test.json")

	// fmt.Println(tasks)

	// fmt.Printf("DB PATH: %v ", ilistlib.DBPath())

	db, err := ilistlib.CreateDB(ilistlib.DBPath())
	checkErr(err)
	checkErr(db.AddUsersTable())
	db.AddUser("snj", "123")
	status, err := db.IsExistsUser("snj", "123")
	checkErr(err)
	println(status)
	status, err = db.IsExistsUser("snj2", "123")
	checkErr(err)
	println(status)

	_, err = db.GetUserByName("snj")
	checkErr(err)

	_, err = db.GetUserByName("snj2")
	checkErr(err)

}
