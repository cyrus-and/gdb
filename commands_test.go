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
		t.Error(result, "!=", expected)
	}

	// gdb-version should generate some console records
	if consoleRecordCount == 0 {
		t.Error("no console records received")
	}

	// try sending a wrong command
	result, err = gdb.Send("foobarbaz")
	if err != nil {
		t.Fatal(err)
	}
	if class, hasClass := result["class"]; !hasClass || class != "error" {
		t.Error(result, "has not class 'error'")
	}

	// exit gdb
	if err := gdb.Exit(); err != nil {
		t.Fatal(err)
	}
}
