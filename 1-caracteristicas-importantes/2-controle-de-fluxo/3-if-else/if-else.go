package main

import "fmt"

func main() {
	exemplo := 9

	if exemplo >= 11 {
		fmt.Println("Maior ou igual a 11")
	} else {
		fmt.Println("Menor que 11")
	}

	exemplo2 := 10

	if exemplo2 == 11 {
		fmt.Println("É igual a 11")
	} else if exemplo2 > 10 {
		fmt.Println("Menor que 10")
	} else {
		fmt.Println("É igual a 10")
	}
}
