package main

import "fmt"

type todolist struct {
	date     string
	Reminder string
	Priority int
}

func (t *todolist) String() string {
	return fmt.Sprintf("%s: %d, %s", t.date, t.Priority, t.Reminder)
}

func (t *todolist) add(value int) {
	t.Priority = value + t.Priority

}
func greaterImportance(first, second todolist) todolist {
	if first.Priority < second.Priority {
		return first
	}
	return second
}

func main() {
	var firstreminder todolist
	firstreminder.date = "06/02/2020"
	firstreminder.Reminder = "Complete Math's Assignment"
	firstreminder.Priority = 2

	var secondreminder todolist
	secondreminder.date = "08/02/2020"
	secondreminder.Reminder = "Complete Assignment Three for CS"
	secondreminder.Priority = 3

	var thirdreminder todolist
	thirdreminder.date = "09/02/2020"
	thirdreminder.Reminder = "Study for Math's midterm!"
	thirdreminder.Priority = 1

	fmt.Printf("%s\n", firstreminder.String())

	firstreminder.add(99)
	fmt.Printf("%s\n", firstreminder.String())
}
