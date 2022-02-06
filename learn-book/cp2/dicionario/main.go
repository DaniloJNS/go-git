// Package main provides ...
package main

import (
	"fmt"
	"os"
	"strings"
)

func colherEstastistica(palavras []string) map[string]int {
	estastisticas := make(map[string]int)

	for _, palavra := range palavras {
		inicial := strings.ToUpper(string(palavra[0]))
		contador, encontrado := estastisticas[inicial]
		if encontrado {
			estastisticas[inicial] = contador + 1
		} else {
			estastisticas[inicial] = 1
		}
	}

	return estastisticas
}

func imprimirEstastistica(estastistica map[string]int) {
	fmt.Println("Contagem de palavras iniciadas em cada letras:")

	for inicial, contador := range estastistica {
		fmt.Printf("%s = %d\n", inicial, contador)
	}

}

func main() {
	palavras := os.Args[1:]

	estastistica := colherEstastistica(palavras)

	imprimirEstastistica(estastistica)
}
