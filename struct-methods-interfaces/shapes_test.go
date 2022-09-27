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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{
			shape: Rectangle{Height: 12, Width: 6},
			want:  72.0,
		},
		{
			shape: Circle{Radius: 10},
			want:  314.1592653589793,
		},
		{
			shape: Triangle{Base: 10, Height: 8},
			want:  40.0,
		},
	}

	for _, testCase := range areaTests {
		got := testCase.shape.Area()
		assert.Equal(t, testCase.want, got)
	}

}
