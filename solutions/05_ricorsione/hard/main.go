package main

import "fmt"

func GenerateParentheses(n int) []string {
	if n <= 0 {
		return []string{}
	}
	out := make([]string, 0)
	var dfs func(open, close int, cur string)
	dfs = func(open, close int, cur string) {
		if len(cur) == 2*n {
			out = append(out, cur)
			return
		}
		if open < n {
			dfs(open+1, close, cur+"(")
		}
		if close < open {
			dfs(open, close+1, cur+")")
		}
	}
	dfs(0, 0, "")
	return out
}

func main() {
	fmt.Println(GenerateParentheses(3))
}
