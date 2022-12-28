package http_server

import (
	"io"
	"testing"
)

// tape_test.go
func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &Tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
