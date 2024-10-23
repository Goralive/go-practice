package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dirMs directMessage) importance() int {
	if dirMs.isUrgent {
		return 50
	}
	return dirMs.priorityLevel
}

func (grMsg groupMessage) importance() int {
	return grMsg.priorityLevel
}

func (smAlert systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch nt := n.(type) {
	case directMessage:
		return nt.senderUsername, nt.importance()
	case groupMessage:
		return nt.groupName, nt.importance()
	case systemAlert:
		return nt.alertCode, nt.importance()
	default:
		return "", 0
	}
}
