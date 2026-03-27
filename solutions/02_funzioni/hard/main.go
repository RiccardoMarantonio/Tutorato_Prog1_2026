package main

import "fmt"

func ComposeAllSafe(x int, funcs ...func(int) (int, bool)) (int, bool) {
	out := x
	for _, fn := range funcs {
		next, ok := fn(out)
		if !ok {
			return 0, false
		}
		out = next
	}
	return out, true
}

func main() {
	res, ok := ComposeAllSafe(2,
		func(v int) (int, bool) { return v + 3, true },
		func(v int) (int, bool) { return v * 2, true },
	)
	fmt.Println(res, ok)
}
