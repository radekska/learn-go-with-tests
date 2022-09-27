package struct_methods_interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Height: 10, Width: 10}
	assert.Equal(t, 40.0, rectangle.Perimeter())
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assert.Equal(t, want, got)
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Height: 12, Width: 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{Radius: 10}
		checkArea(t, circle, 314.1592653589793)
	})
}
