package gdb

import (
	"fmt"
	"testing"
)

func TestIO(t *testing.T) {
	input := "package   foo"
	expectedIn := fmt.Sprintf("%s\r\n", input) // not sure why \r\n...
	expectedOut := "package foo\r\n"

	// start gdb
	gdb, err := New(nil)
	if err != nil {
		t.Fatal(err)
	}

	// start processing the output
	done := make(chan bool)
	go func() {
		var n int
		buf := make([]byte, 1024)

		// read the first line
		n, err = gdb.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		if string(buf[:n]) != expectedIn {
			fmt.Printf("'%s'\n", buf[:n])
			fmt.Printf("'%s'\n", []byte(expectedIn))
			t.Fatal("should read back the input")
		}

		// read the second line
		n, err = gdb.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		if string(buf[:n]) != expectedOut {
			fmt.Printf("'%s'\n", buf[:n])
			fmt.Printf("'%s'\n", []byte(expectedOut))
			t.Fatal("should read the proper output")
		}

		// try another read
		n, err = gdb.Read(buf)
		if err == nil {
			t.Fatal("read should fail")
		}

		done <- true
	}()

	// load a program
	if _, err = gdb.Send("file-exec-file", "gofmt"); err != nil {
		t.Fatal(err)
	}

	// provide some input
	if _, err := gdb.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if _, err := gdb.Write([]byte("\n\x04")); err != nil {
		t.Fatal(err)
	}

	// run the program
	if _, err = gdb.Send("exec-run"); err != nil {
		t.Fatal(err)
	}

	// exit gdb
	if err := gdb.Exit(); err != nil {
		t.Fatal(err)
	}

	// wait for the output processing
	_ = <-done
}
