package main

import (
	"fmt"
)

func main() {
	var classe, familia, genero, especie string

	fmt.Scan(&classe, &familia, &genero)

	switch genero {
	case "carnivoro":
		especie = "aguia"
	case "onivoro":
		switch familia {
		case "ave":
			especie = "pomba"
		case "mamifero":
			especie = "homem"
		case "anelideo":
			especie = "minhoca"
		}
	case "herbivoro":
		switch familia {
		case "inseto":
			especie = "lagarta"
		case "mamifero":
			especie = "vaca"
		}
	case "hematofago":
		switch familia {
		case "inseto":
			especie = "pulga"
		case "anelideo":
			especie = "sanguessuga"
		}
	}
	fmt.Println(especie)
}
