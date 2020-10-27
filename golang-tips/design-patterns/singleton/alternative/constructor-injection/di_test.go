package di

func ExampleLogger() {
	userLogger := NewLogger("user")
	contentLogger := NewLogger("content")

	user := NewUser(userLogger)
	user.Print()
	content := NewContent(contentLogger)
	content.Print()

	// Output:
	// user: test user print
	// content: test content print
}

func ExampleSingletonLogger() {
	logger := NewLogger("app")
	user := NewUser(logger)
	user.Print()
	content := NewContent(logger)
	content.Print()

	// Output:
	// app: test user print
	// app: test content print
}
