package gdb

import (
	"testing"
	"time"
)

func TestInterrupt(t *testing.T) {
	// start gdb
	gdb, err := New(nil)
	if err != nil {
		t.Fatal(err)
	}

	// load and start gofmt (a program which will wait for input)
	_, err = gdb.Send("file-exec-and-symbols", "gofmt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = gdb.Send("exec-run")
	if err != nil {
		t.Fatal(err)
	}

	interrupt := false
	go func() {
		// allow the program to start up
		time.Sleep(time.Second)

		// then interrupt the execution
		interrupt = true
		if err := gdb.Interrupt(); err != nil {
			t.Fatal(err)
		}
	}()

	// try to execute a command
	_, err = gdb.Send("gdb-version")
	if err != nil {
		t.Fatal(err)
	}
	if interrupt == false {
		t.Fatal("GDB should wait for the interrupt")
	}

	// exit gdb
	if err := gdb.Exit(); err != nil {
		t.Fatal(err)
	}
}
