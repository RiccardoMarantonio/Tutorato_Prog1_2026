package main

import "testing"

func TestContaParole(t *testing.T) {
	tests := []struct {
		name   string
		parole []string
		want   map[string]int
	}{
		{
			"base",
			[]string{"ciao", "mondo", "ciao", "go"},
			map[string]int{"ciao": 2, "mondo": 1, "go": 1},
		},
		{
			"vuoto",
			[]string{},
			map[string]int{},
		},
		{
			"tutte uguali",
			[]string{"go", "go", "go"},
			map[string]int{"go": 3},
		},
		{
			"tutte diverse",
			[]string{"a", "b", "c"},
			map[string]int{"a": 1, "b": 1, "c": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContaParole(tt.parole)
			if len(got) != len(tt.want) {
				t.Errorf("ContaParole(%v) size = %d, want %d", tt.parole, len(got), len(tt.want))
				return
			}
			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("ContaParole(%v)[%q] = %d, want %d", tt.parole, k, got[k], v)
				}
			}
		})
	}
}
