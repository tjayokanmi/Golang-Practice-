package main

import "fmt"

func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"
	case "mac":
		creator = "A steve"
	default:
		creator = "unknown"
	}
	return creator
}

func main() {
	fmt.Println(getCreator("linux"))
	fmt.Println(getCreator("windows"))
	fmt.Println(getCreator("mac"))
}
