package singleton

import (
	"fmt"
	"sync"
)

var singleLogger *Logger
var once sync.Once

type User struct {
	logger *Logger
}

func (u *User) Print() {
	u.logger.Print("test user print")
}

func NewUser() *User {
	return &User{logger: GetLogger("user")}
}

type Content struct {
	logger *Logger
}

func (u *Content) Print() {
	u.logger.Print("test content print")
}

func NewContent() *Content {
	return &Content{logger: GetLogger("content")}
}

type Logger struct {
	appName string
}

func (l *Logger) Print(msg string) {
	fmt.Printf("%s: %s\n", l.appName, msg)
}

func GetLogger(name string) *Logger {
	once.Do(func() {
		singleLogger = &Logger{
			appName: name,
		}
	})
	return singleLogger
}
