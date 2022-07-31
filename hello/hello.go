package hello

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Tri", ""))
}

const indonesia = "Indonesia"
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const indonesiaHelloPrefix = "Hai, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case indonesia:
		prefix = indonesiaHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
