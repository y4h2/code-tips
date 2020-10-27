package singleton

func ExampleLogger() {
	user := NewUser()
	user.Print()

	content := NewContent()
	content.Print()
	// Output:
	// user: test user print
	// content: test content print
}
