package main

import "fmt"

// alias
type Filter func(string) string

func sayHelloFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spanFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("Mamah", spanFilter)

	filter := spanFilter

	sayHelloFilter("Anjing", filter)
}
