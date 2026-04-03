package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func ContaParole(parole []string) map[string]int {
	contatori := make(map[string]int)
	for _, p := range parole {
		contatori[p]++
	}
	return contatori
}

func main() {
	var parole []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parole = append(parole, line)
	}

	contatori := ContaParole(parole)

	var chiavi []string
	for k := range contatori {
		chiavi = append(chiavi, k)
	}
	sort.Strings(chiavi)

	for _, k := range chiavi {
		fmt.Printf("%s: %d\n", k, contatori[k])
	}
}
