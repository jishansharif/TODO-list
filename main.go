package main

import "fmt"

type Item struct {
	date   string
	detail string
}

func (i *Item) String() string {
	return fmt.Sprintf("%s: %s", i.date, i.detail)
}

type List struct {
	items []Item
}

func (l *List) Add(i Item) {
	l.items = append(l.items, i)
}

func (l *List) Remove(toRemove Item) {
	foundIdx := -1
	for i := range l.items {
		if l.items[i] == toRemove {
			foundIdx = i
			break
		}
	}

	if foundIdx == -1 {
		// We didn't find anything
		return
	}

	foundBefore := l.items

	l.items = foundBefore[:foundIdx]
	l.items = append(l.items, foundBefore[foundIdx+1:]...)
}

func (l *List) String() string {
	var result string
	for _, item := range l.items {
		result += fmt.Sprintf(item.String())
		result += "\n"
	}

	return result
}

func x() {
	// TODO(jishansharif): Complete this application.
	// Each loop, ask if user wants to do (1) or (2).
	//
	// 1. Get data + detail from user
	//  1(a). Save list to disk
	// 2. Show me current list
	//
	// l.Save("todo.list")
	// n.Load("todo.list")

	var l List
	jrent := Item{
		date:   "1/2/2020",
		detail: "Pay Jishan's rent",
	}

	irent := Item{
		date:   "1/2/2020",
		detail: "Pay Irfan's rent",
	}

	l.Add(jrent)
	l.Add(irent)

	fmt.Printf("TODO List\n\n%s", l.String())
}
