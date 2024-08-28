package main

import "fmt"

func splitEmail(email string) (string, string) {

	username, domain := "", ""

	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

func main() {
	username, domain := splitEmail("john@gmail.com")
	fmt.Println(username + domain)
}
