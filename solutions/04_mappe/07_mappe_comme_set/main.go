package main

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func Add(s Set, elemento string) {
	s[elemento] = struct{}{}
}

func Remove(s Set, elemento string) {
	delete(s, elemento)
}

func Contains(s Set, elemento string) bool {
	_, exists := s[elemento]
	return exists
}

func Union(s1, s2 Set) Set {
	result := NewSet()
	for k := range s1 {
		Add(result, k)
	}
	for k := range s2 {
		Add(result, k)
	}
	return result
}

func Intersection(s1, s2 Set) Set {
	result := NewSet()
	for k := range s1 {
		if Contains(s2, k) {
			Add(result, k)
		}
	}
	return result
}

func Size(s Set) int {
	return len(s)
}

func main() {}
