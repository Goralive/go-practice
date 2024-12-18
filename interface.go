package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed {
		return len(e.body) * 2
	}
	return len(e.body) * 5
}

func (e email) format() string {
	var sub string
	if e.isSubscribed {
		sub = "Subscribed"
	} else {
		sub = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, sub)

}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

