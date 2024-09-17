package main

import "testing"

// Writing a test is just like writing a function, with a few rules
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// To use the *testing.T type, you need to import "testing", like we did with fmt in the other file
// For now, it's enough to know that your t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail()
// when you want to fail.
// We've covered some new topics:
// if
// If statements in Go are very much like other programming languages.
// Declaring variables
// We're declaring some variables with the syntax varName := value, which lets us reuse some values in our test for readability.
// t.Errorf
// We are calling the Errorf method on our t, which will print out a message and fail the test. The f stands for format, which allows us to build a string with values inserted into the placeholder values %q. When you make the test fail, it should be clear how it works.

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
