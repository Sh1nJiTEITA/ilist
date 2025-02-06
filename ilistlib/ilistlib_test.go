package ilistlib

import "testing"

func testTask(t *testing.T) {
	msg := "DO IT"
	sts := false

	task := Task{sts, msg}

	if task.Content != msg || task.Status != sts {
		t.Fatalf("Incorrect Task structure constructor")
	}
}
