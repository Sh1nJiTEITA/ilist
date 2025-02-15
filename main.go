package main

import (
	// "fmt"
	"fmt"
	"os"

	"github.com/Sh1nJiTEITA/ilist/ilistlib"
	"github.com/Sh1nJiTEITA/ilist/interaction"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
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
	checkErr(db.AddUserTable())
	checkErr(db.AddTaskTable())

	// user, err := db.AddUser("snj", "123")
	// checkErr(err)

	status, err := db.IsExistsUser("snj", "123")
	checkErr(err)
	println(status)
	status, err = db.IsExistsUser("snj2", "123")
	checkErr(err)
	println(status)

	user, err := db.GetUserByName("snj")
	checkErr(err)

	_, err = db.AddTask(user, "do something by snj", false, 0)

	cli, err := interaction.CreateCLIManager(db)
	checkErr(err)

	msg, err := cli.GetTasksByUserStr(user.Id)
	checkErr(err)

	fmt.Print(msg)

}
