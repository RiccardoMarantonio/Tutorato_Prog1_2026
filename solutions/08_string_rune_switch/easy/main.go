package main

import "fmt"

func FirstRune(s string) (rune, bool) {
	for _, r := range s {
		return r, true
	}
	return 0, false
}

func main() {
	r, ok := FirstRune("ciao")
	fmt.Println(r, ok)
}
