package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Decrypt(s string, k int) string {
	runes := []rune(s)
	k = k % 26
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = rune((int(runes[i]-'a')-k+26)%26 + 'a')
		}
	}
	return string(runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	text := strings.Join(lines, "\n")

	decifrato := Decrypt(text, k)

	fmt.Printf("Decifrato: %s\n", decifrato)
}
