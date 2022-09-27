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
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle",
			shape: Rectangle{Height: 12, Width: 6},
			want:  72.0,
		},
		{
			name:  "Circle",
			shape: Circle{Radius: 10},
			want:  314.1592653589793,
		},
		{
			name:  "Triangle",
			shape: Triangle{Base: 10, Height: 8},
			want:  40.0,
		},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			assert.Equal(t, testCase.want, got)
		})
	}

}
