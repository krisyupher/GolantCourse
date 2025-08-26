package main

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
func testMain() {
	var greetingRespont = Hello("Glad6ys")
	fmt.Println(Hello(greetingRespont))
}
