package iteration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	assert.Equal(t, "aaaaa", Repeat("a", 5))
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 10)
	fmt.Println(repeated)
	// Output: bbbbbbbbbb
}
