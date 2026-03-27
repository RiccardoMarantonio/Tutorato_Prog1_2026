package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func RunProgram(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	acc := 0
	ended := false

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) == 1 && parts[0] == "end" {
			ended = true
			break
		}
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid command format")
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		switch parts[0] {
		case "set":
			acc = n
		case "add":
			acc += n
		case "sub":
			acc -= n
		case "mul":
			acc *= n
		default:
			return 0, fmt.Errorf("invalid command")
		}
	}
	if err := s.Err(); err != nil {
		return 0, err
	}
	if !ended {
		return 0, fmt.Errorf("missing end")
	}
	return acc, nil
}

func main() {
	v, err := RunProgram(strings.NewReader("set 5\nadd 1\nend\n"))
	fmt.Println(v, err)
}
