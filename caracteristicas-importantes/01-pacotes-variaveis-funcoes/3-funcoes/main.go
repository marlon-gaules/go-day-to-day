package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return a, b
}

func main() {
	fmt.Println(add(13, 13))

	c, d := swap("Hello", "world!")
	fmt.Println(c, d)
}
