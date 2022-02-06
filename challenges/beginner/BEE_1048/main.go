package main

import "fmt"

func main() {
	var salario float32
	var percentual int

	fmt.Scanf("%f", &salario)

	if salario >= 0 && salario <= 400.00 {
		percentual = 15
	} else if salario > 400.00 && salario <= 800.00 {
		percentual = 12
	} else if salario > 800.00 && salario <= 1200.00 {
		percentual = 10
	} else if salario > 1200.00 && salario <= 2000.00 {
		percentual = 7
	} else if salario > 2000.00 {
		percentual = 4
	}

	reajuste := salario * float32(percentual) / 100
	novo_salario := salario + reajuste

	fmt.Printf("Novo salario: %.2f\n", novo_salario)
	fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
	fmt.Println("Em percentual:", percentual, "%")
}
