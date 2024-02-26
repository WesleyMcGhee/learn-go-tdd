package main

import "fmt"

func hello() string {
	return "Hello, world"
}

func helloName(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(hello())
}