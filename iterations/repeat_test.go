package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, want, got string) {
		t.Helper()

		if got != want {
			t.Errorf("expected %q but got %q", got, want)
		}
	}

	t.Run("When a character is provided with count", func(t *testing.T) {
		repeated := Repeat("a", 4)
		expected := "aaaa"

		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("When no character is provided returns empty string", func(t *testing.T) {
		repeated := Repeat("", 1)
		expected := ""

		assertCorrectMessage(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
