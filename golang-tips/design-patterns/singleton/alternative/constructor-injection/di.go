package di

import "fmt"

type User struct {
	logger *Logger
}

func (u *User) Print() {
	u.logger.Print("test user print")
}

func NewUser(logger *Logger) *User {
	return &User{logger: logger}
}

type Content struct {
	logger *Logger
}

func (u *Content) Print() {
	u.logger.Print("test content print")
}

func NewContent(logger *Logger) *Content {
	return &Content{logger: logger}
}

type Logger struct {
	appName string
}

func (l *Logger) Print(msg string) {
	fmt.Printf("%s: %s\n", l.appName, msg)
}

func NewLogger(name string) *Logger {
	return &Logger{
		appName: name,
	}
}
