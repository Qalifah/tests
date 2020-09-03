package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	item := "i can do this shit! "
	get := Repeat(item, 6)
	want := item + item + item + item + item + item

	if get != want {
		t.Errorf("wanted %q but got %q", want, get)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Repeat("omg", 10)
	}
}

func ExampleRepeat() {
	ans := Repeat("ba", 3)
	fmt.Printf(`"%s"`, ans)
	// Output: "bababa"
}