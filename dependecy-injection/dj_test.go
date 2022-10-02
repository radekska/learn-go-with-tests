package dependecy_injection

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	assert.Equal(t, want, got)
}
