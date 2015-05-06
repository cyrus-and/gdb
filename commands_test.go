package gdb

import (
	"reflect"
	"testing"
)

var consoleRecordCount = 0

func onNotification(t *testing.T, notification map[string]interface{}) {
	if notification["type"] != "console" {
		consoleRecordCount++
	}
}

func TestSend(t *testing.T) {
	var expected, result map[string]interface{}

	// start gdb
	gdb, err := New(func(notification map[string]interface{}) {
		onNotification(t, notification)
	})
	if err != nil {
		t.Fatal(err)
	}

	// try sending a command
	result, err = gdb.Send("gdb-version")
	expected = map[string]interface{}{"class": "done"}
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatal(result, "!=", expected)
	}

	// gdb-version should generate some console records
	if consoleRecordCount == 0 {
		t.Fatal("no console records received")
	}

	// try sending a command with an argument containing a space
	result, err = gdb.Send("var-create", "foo", "@", "40 + 2")
	if err != nil {
		t.Fatal(err)
	}
	if class, hasClass := result["class"]; !hasClass || class != "done" {
		t.Fatal(result, "has not class 'done'")
	}

	// try sending a wrong command
	result, err = gdb.Send("foobarbaz")
	if err != nil {
		t.Fatal(err)
	}
	if class, hasClass := result["class"]; !hasClass || class != "error" {
		t.Fatal(result, "has not class 'error'")
	}

	// exit gdb
	if err := gdb.Exit(); err != nil {
		t.Fatal(err)
	}
}
