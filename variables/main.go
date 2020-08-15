package main

import "fmt"

func main() {
	// var dog, cat string = "🐕", "🐈"
	dog, cat := "🐕", "🐈"
	cat = "😾"
	cat, face := "gato", "😵"

	fmt.Println(dog, cat, face)
}
