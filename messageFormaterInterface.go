package main

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func (plainText PlainText) Format() string {
	return plainText.message
}

func (bold Bold) Format() string {
	boldedMessage := "**" + bold.message + "**"
	return boldedMessage
}

func (code Code) Format() string {
	codeFragment := "`" + code.message + "`"
	return codeFragment
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
