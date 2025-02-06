package ilistlib

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Status  bool   `json:"status"`
	Content string `json:"content"`
}

func (t Task) String() string {
	return fmt.Sprintf("Task(Status: %v, Content: %s)", t.Status, t.Content)
}

func WriteTaskToFile(task *Task, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	row := fmt.Sprintf("{\"status\": %v, \"content\": \"%v\"}", task.Status, task.Content)
	file.WriteString(row)
}

func WriteTasksToFile(data *[]Task, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	jsonData, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		panic(err)
	}
	file.Write(jsonData)
}

func ReadTasksFromFile(data *[]Task, path string) {
	if _, err := os.Stat(path); err != nil {
		panic(err)
	}
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}
}
