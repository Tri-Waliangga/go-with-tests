package iteration

import (
	"fmt"
	"testing"
)

const repeatTimesChar = 9
const repeatedChar = "a"

func TestRepeat(t *testing.T) {

	repeated := Repeat(repeatedChar, repeatTimesChar)
	var expected string

	for i := 0; i < repeatTimesChar; i++ {
		expected += repeatedChar
	}

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(repeatedChar, repeatTimesChar)
	}
}

func ExampleRepeat() {
	charRepeated := Repeat("a", 10)
	fmt.Println(charRepeated)
	// Output: "aaaaaaaaaa"
}
