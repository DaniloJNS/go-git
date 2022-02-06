package main

import (
	"fmt"
)

func main() {
	var input, input2, resultado int

	fmt.Scanf("%d %d", &input, &input2)

	if input > input2 {
		resultado = 24 + input2 - input
	} else if input == input2 {
		resultado = 24
	} else {
		resultado = input2 - input
	}

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", resultado)
}
