package shape

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Height: 12.0, Width: 6.0}
		got := rectangle.Area()
		want := 72.0
		assertCorrectMessage(t, got, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
