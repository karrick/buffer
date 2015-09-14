package buffer

import (
	"bytes"
	"io"
	"testing"
)

func TestReadEmpty(t *testing.T) {
	b := Buffer{}
	buf := make([]byte, 4096)
	n, err := b.Read(buf)
	if err != io.EOF {
		t.Errorf("Actual: %#v; Expected: %#v", err, io.EOF)
	}
	if n != 0 {
		t.Errorf("Actual: %#v; Expected: %#v", n, 0)
	}
}

func TestWriteEmpty(t *testing.T) {
	b := Buffer{}
	buf := []byte("this is a test")
	n, err := b.Write(buf)
	if err != nil {
		t.Errorf("Actual: %#v; Expected: %#v", err, nil)
	}
	if want := len(buf); n != want {
		t.Errorf("Actual: %#v; Expected: %#v", n, want)
	}
	if !bytes.Equal(buf, b) {
		t.Errorf("Actual: %#v; Expected: %#v", buf, b)
	}
}
