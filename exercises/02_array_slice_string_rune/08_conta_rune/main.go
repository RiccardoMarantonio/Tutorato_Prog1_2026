package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Conteggio struct {
	Totale  int
	Lettere int
	Cifre   int
	Spazi   int
	Altri   int
}

func ContaRune(s string) Conteggio {
	c := Conteggio{}
	for _, r := range s {
		c.Totale++
		if unicode.IsLetter(r) {
			c.Lettere++
		} else if unicode.IsDigit(r) {
			c.Cifre++
		} else if unicode.IsSpace(r) {
			c.Spazi++
		} else {
			c.Altri++
		}
	}
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	c := ContaRune(input)

	fmt.Printf("Totale rune: %d\n", c.Totale)
	fmt.Printf("Lettere: %d\n", c.Lettere)
	fmt.Printf("Cifre: %d\n", c.Cifre)
	fmt.Printf("Spazi: %d\n", c.Spazi)
	fmt.Printf("Altri: %d\n", c.Altri)
}
