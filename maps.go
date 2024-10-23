package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	users := make(map[string]user)

	if len(names) != len(phoneNumbers) {
		return users, errors.New("invalid sizes")
	}

	for index, name := range names {
		users[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[index],
		}
	}

	return users, nil
}

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, name := range messagedUsers {
		if _, ok := validUsers[name]; ok {
			validUsers[name] += 1
		}
	}
}

type user struct {
	name        string
	phoneNumber int
}
