package main

import "testing"

func TestSetBasics(t *testing.T) {
	s := NewSet()

	Add(s, "a")
	if !Contains(s, "a") {
		t.Error("dopo Add, Contains dovrebbe restituire true")
	}
	if Contains(s, "b") {
		t.Error("Contains su elemento non aggiunto dovrebbe restituire false")
	}

	if Size(s) != 1 {
		t.Errorf("Size dopo 1 Add = %d, want 1", Size(s))
	}
}

func TestSetRemove(t *testing.T) {
	s := NewSet()
	Add(s, "a")
	Add(s, "b")

	Remove(s, "a")

	if Contains(s, "a") {
		t.Error("dopo Remove, Contains dovrebbe restituire false")
	}
	if !Contains(s, "b") {
		t.Error("Remove non dovrebbe rimuovere altri elementi")
	}

	Remove(s, "inexistent")

	if Size(s) != 1 {
		t.Errorf("Size dopo Remove su elemento inesistente = %d, want 1", Size(s))
	}
}

func TestSetUnion(t *testing.T) {
	s1 := NewSet()
	Add(s1, "a")
	Add(s1, "b")

	s2 := NewSet()
	Add(s2, "b")
	Add(s2, "c")

	unione := Union(s1, s2)

	if !Contains(unione, "a") {
		t.Error("Union dovrebbe contenere elementi di s1")
	}
	if !Contains(unione, "b") {
		t.Error("Union dovrebbe contenere elementi comuni")
	}
	if !Contains(unione, "c") {
		t.Error("Union dovrebbe contenere elementi di s2")
	}
	if Size(unione) != 3 {
		t.Errorf("Union size = %d, want 3", Size(unione))
	}
}

func TestSetIntersection(t *testing.T) {
	s1 := NewSet()
	Add(s1, "a")
	Add(s1, "b")
	Add(s1, "c")

	s2 := NewSet()
	Add(s2, "b")
	Add(s2, "c")
	Add(s2, "d")

	inter := Intersection(s1, s2)

	if Contains(inter, "a") {
		t.Error("Intersection non dovrebbe contenere elementi solo in s1")
	}
	if !Contains(inter, "b") {
		t.Error("Intersection dovrebbe contenere elementi comuni")
	}
	if !Contains(inter, "c") {
		t.Error("Intersection dovrebbe contenere elementi comuni")
	}
	if Contains(inter, "d") {
		t.Error("Intersection non dovrebbe contenere elementi solo in s2")
	}
	if Size(inter) != 2 {
		t.Errorf("Intersection size = %d, want 2", Size(inter))
	}
}

func TestSetNoDuplicates(t *testing.T) {
	s := NewSet()
	Add(s, "a")
	Add(s, "a")
	Add(s, "a")

	if Size(s) != 1 {
		t.Errorf("Add duplicato non dovrebbe aumentare la size: got %d, want 1", Size(s))
	}
}

func TestSetEmpty(t *testing.T) {
	s := NewSet()

	if Size(s) != 0 {
		t.Errorf("NewSet dovrebbe essere vuoto: got %d, want 0", Size(s))
	}

	unione := Union(s, s)
	if Size(unione) != 0 {
		t.Errorf("Union di set vuoti dovrebbe essere vuoto: got %d, want 0", Size(unione))
	}

	inter := Intersection(s, s)
	if Size(inter) != 0 {
		t.Errorf("Intersection di set vuoti dovrebbe essere vuoto: got %d, want 0", Size(inter))
	}
}
