package main

import (
	"fmt"
)

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
	var afterFiltermessages []Message
	for _, msg := range messages {
		if msg.Type() == filterType {
			afterFiltermessages = append(afterFiltermessages, msg)
		}
	}
	return afterFiltermessages
}

func isValidPassword(password string) bool {

	hasUpperCase := false
	hasDitgit := false

	var pwLength int = len(password)
	if pwLength < 5 || 12 < pwLength {
		return false
	}
	for _, letter := range password {
		if letter >= 'A' && letter <= 'Z' {
			hasUpperCase = true
		}
		if letter >= '0' && letter <= '9' {
			hasDitgit = true
		}

	}

	return hasDitgit && hasUpperCase

}

func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func(s string) int {
			return len(s) * 2
		}, message)

	}
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
