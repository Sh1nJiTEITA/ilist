package main

import (
	// "fmt"
	// "fmt"
	"os"

	mod_lib "github.com/Sh1nJiTEITA/ilist/ilistlib"
	mod_int "github.com/Sh1nJiTEITA/ilist/interaction"
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
	db, err := mod_lib.OpenDB("./main.db")
	if err != nil {
		panic(err)
	}

	users := mod_lib.NewTableUsers(db)
	if !db.IsTableExists(users) {
		db.CreateTable(users)
	}
	if user, err := users.FindByName("snj"); err != nil {
		user = mod_lib.NewUser("snj", "123")
		err = users.Save(user)
		if err != nil {
			panic(err)
		}
	}

	tables := []mod_int.CliTableCommand{
		users,
		mod_int.TestCommand{},
	}

	mod_int.ParseInputArguments(tables, os.Args)
}
