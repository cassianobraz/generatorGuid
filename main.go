package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
	for {
		var quantidade int

		fmt.Print("Quantos GUIDs deseja gerar? ")

		if _, err := fmt.Scanln(&quantidade); err != nil || quantidade <= 0 {
			fmt.Println("Digite um número inteiro maior que 0.\n")
			limparEntrada()
			continue
		}

		fmt.Println()

		for i := 1; i <= quantidade; i++ {
			fmt.Printf("%2d => %s\n", i, uuid.New())
		}

		var opcao string

		fmt.Print("\nDeseja gerar mais GUIDs? (S/N): ")
		fmt.Scanln(&opcao)

		if strings.EqualFold(opcao, "N") {
			fmt.Println("\nPrograma finalizado.")
			break
		}

		fmt.Println()
	}
}

func limparEntrada() {
	var lixo string
	fmt.Scanln(&lixo)
}
