package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestUsageErr(t *testing.T) {
	ok, err := Usage(nil)

	if err == nil {
		t.Fail()
	}

	if ok {
		t.Fail()
	}
}

func TestUsageOK(t *testing.T) {
	var ok bool

	out := captureStdout(func() {
		ok, _ = Usage(os.Stdin)
	})

	if !ok {
		t.Fail()
	}

	if len(out) == 0 {
		t.Fail()
	}
}

func TestUsageNotOK(t *testing.T) {
	f, _ := ioutil.TempFile("", "test")
	defer f.Close()

	f.Write([]byte("FOO=foo"))

	if ok, _ := Usage(f); ok {
		t.Fail()
	}
}

//
// Helper
//

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}