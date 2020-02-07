package main

type todolist struct {
	date     string
	Reminder string
	Priority int
}

func greaterimportance(first, second, todolist) todolist {
	if first.Priority < second.Priority {
		return first
	}
	return second
}

func main() {
	firstreminder := todolist
	date := "06/02/2020"
	Reminder := "Complete Math's Assignment"
	Priority := 2

	secondreminder := firstreminder
	secondreminder.date = "08/02/2020"
	secondreminder.Reminder = "Complete Assignment Three for CS"
	secondreminder.Priority = 3

	thirdreminder := firstreminder
	thirdreminder.date = "09/02/2020"
	thirdreminder.Reminder = "Study for Math's midterm!"
	thirdreminder.Priority = 1

}
