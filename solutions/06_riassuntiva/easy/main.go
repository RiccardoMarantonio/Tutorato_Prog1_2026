package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

func BuildPhonebook(contacts []Contact) map[string]string {
	book := make(map[string]string)
	for _, c := range contacts {
		book[c.Name] = c.Phone
	}
	return book
}

func main() {
	items := []Contact{{Name: "Anna", Phone: "123"}, {Name: "Anna", Phone: "456"}}
	fmt.Println(BuildPhonebook(items))
}
