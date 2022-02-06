package main

import "fmt"

func main() {
	var Hi, Mi, Hf, Mf, resultado int
	fmt.Scanf("%d %d %d %d", &Hi, &Mi, &Hf, &Mf)

	timestamp_inicial := Hi*60 + Mi
	timestamp_final := Hf*60 + Mf

	if timestamp_final > timestamp_inicial {
		resultado = timestamp_final - timestamp_inicial
	} else if timestamp_inicial == timestamp_final {
		resultado = 1440
	} else {
		resultado = 1440 + timestamp_final - timestamp_inicial
	}

	resultado_hora := resultado / 60
	resultado_minuto := resultado - resultado_hora*60

	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", resultado_hora, resultado_minuto)
}
