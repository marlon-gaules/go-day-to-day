package main

import "fmt"

func main() {
	diaDaSemana := "Sabado"

	switch diaDaSemana {
	case "Domingo":
		fmt.Println("Hoje é domingo")
	case "Segunda":
		fmt.Println("Hoje é segunda-feira")
	case "Terça":
		fmt.Println("Hoje é terça-feira")
	case "Quarta":
		fmt.Println("Hoje é quarta-feira")
	case "Quinta":
		fmt.Println("Hoje é quinta-feira")
	case "Sexta":
		fmt.Println("Hoje é sexta-feira")
	default:
		fmt.Println("Hoje é sabado")
	}
}
