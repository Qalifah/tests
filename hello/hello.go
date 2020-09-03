package main

import (
	"fmt"
)

const (
	french = "French"
	frenchHelloPrefix = "Bonjour, "
	spanish = "Spanish"
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Welcome, "
	empty = ""
)

func hello(name, language string) string {
	if name == empty {
		name = "Human"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(hello("Aliens", empty))
}