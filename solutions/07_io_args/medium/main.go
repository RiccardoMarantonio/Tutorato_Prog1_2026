package main

import (
	"fmt"
	"strconv"
)

func ParseCommand(args []string) (int, error) {
	if len(args) != 4 {
		return 0, fmt.Errorf("invalid args length")
	}
	a, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, err
	}
	b, err := strconv.Atoi(args[3])
	if err != nil {
		return 0, err
	}

	switch args[1] {
	case "add":
		return a + b, nil
	case "mul":
		return a * b, nil
	default:
		return 0, fmt.Errorf("invalid command")
	}
}

func main() {
	v, err := ParseCommand([]string{"prog", "add", "2", "5"})
	fmt.Println(v, err)
}
