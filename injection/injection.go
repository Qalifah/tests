package injection

import (
	"io"
	"fmt"
)

func Greet(output io.Writer, words string) {
	fmt.Fprintf(output, "Hello, %s", words)
}