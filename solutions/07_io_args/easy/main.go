package main

import (
	"fmt"
	"io"
	"strings"
)

func SumTwoFromReader(r io.Reader) (int, error) {
	var a, b int
	if _, err := fmt.Fscan(r, &a, &b); err != nil {
		return 0, err
	}
	return a + b, nil
}

func main() {
	v, err := SumTwoFromReader(strings.NewReader("3 4"))
	fmt.Println(v, err)
}
