package structandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("testing perimeter", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := Perimeter(rect)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("testing area of rectangle", func(t *testing.T) {
		rect := Rectangle{Width: 12.0, Height: 6.0}
		checkArea(t, rect, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
