package main

import "fmt"

type Greeting func(name string) string

func say(g Greeting, n string) { fmt.Println(g(n)) }

func french(name string) string { return "Bonjour, " + name }

func main() {
	english := func(name string) string { return "Hello, " + name }

	say(english, "ANisus")
	say(french, "ANisus")
}
