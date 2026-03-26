package greeting

import (
	"fmt"
	"unicode"
)

var Greet = greet

func greet(name ...string) string {
	if len(name) == 0 {
		return fmt.Sprintf("Hello, my friend.")
	}

	r := []rune(name[0])
	isUpper := true
	for _, v := range r {
		if unicode.IsLower(v) {
			isUpper = false
		}
	}

	if isUpper {
		return fmt.Sprintf("HELLO, %s.", name[0])
	}

	return fmt.Sprintf("Hello, %s.", name[0])
}
