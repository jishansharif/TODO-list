package main

import "fmt"

type todolist struct {
	date     string
	Reminder string
}

func main() {
	firstreminder := todolist
	date := "06/02/2020"
	Reminder := "Complete Math's Assignment"

	fmt.Printf("%s: %s\n", date, Reminder)
}
