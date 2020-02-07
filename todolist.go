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

	secondreminder := firstreminder
	secondreminder.date = "08/02/2020"
	secondreminder.Reminder = "Complete Assignment Three for CS"

	thirdreminder := firstreminder
	thirdreminder.date = "09/02/2020"
	thirdreminder.Reminder = "Study for Math's midterm!"

	fmt.Printf("%s: %s\n", date, Reminder)
}
