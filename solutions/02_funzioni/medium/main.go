package main

import "fmt"

func ComposeAll(x int, funcs ...func(int) int) int {
	out := x
	for _, fn := range funcs {
		out = fn(out)
	}
	return out
}

func main() {
	res := ComposeAll(2,
		func(v int) int { return v + 3 },
		func(v int) int { return v * 2 },
	)
	fmt.Println(res)
}
