package main

import "fmt"

func main() {
	for i := 0; i <= 16; i++ {
		fmt.Println(i)
	}

	sum := 0
	for x := 0; x < 10; x++ {
		sum += x
	}
	fmt.Println(sum)
}
