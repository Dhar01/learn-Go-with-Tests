package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("specific counts", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("if count is zero", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "a"
		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Repeat("a", 8)
	}
}
