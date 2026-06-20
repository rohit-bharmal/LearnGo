package main

import "fmt"

func main() {
	// colors is a map of string keys and string values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println(colors)
	printMap(colors)
}

// printMap is a function that prints the hex code for each color
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

/* maps are unordered collections of key-value pairs
maps are indexed by keys, which are unique and immutable
maps are mutable, which means you can add, remove, and modify elements
maps are dynamic, which means you can add and remove keys at any time
maps are unordered, which means you cannot predict the order of the keys
maps are not thread-safe, which means you cannot use maps in concurrent operations
maps are not safe for concurrent operations, which means you cannot use maps in concurrent operations
*/
