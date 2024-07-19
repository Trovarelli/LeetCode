package main

import (
	"fmt"
	"leetcode/desafios"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, especifique o número do desafio a ser executado")
		os.Exit(1)
	}

	desafio := os.Args[1]

	switch desafio {
	case "1":
		desafios.Desafio01()
	case "2":
		desafios.Desafio02()
	case "3":
		desafios.Desafio03()
	default:
		fmt.Println("Desafio não encontrado")
	}
}
