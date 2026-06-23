package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Conversão de inteiro para float
	var valor int = 42
	var valorConvertido float64 = float64(valor)

	fmt.Println("Inteiro:", valor)
	fmt.Println("Float:", valorConvertido)

	// Conversão de inteiro para string
	var valorString string = strconv.Itoa(valor)
	fmt.Println("String:", valorString)

	// Conversão de string para inteiro

	var valorString2 string = "230.4"
	valorInteiro, err := strconv.Atoi(valorString2)

	if err != nil {
		fmt.Println("ERRO AO CONVERTER")
	} else {
		fmt.Println("String convertida em Inteiro:", valorInteiro)
	}

	// Conversão de string para float
	var piString string = "3.14159"
	piConvertido, err2 := strconv.ParseFloat(piString, 64)

	if err2 != nil {
		fmt.Println("ERRO AO CONVERTER PARA FLOAT")
	} else {
		fmt.Println("String convertida em float:", piConvertido)
	}

}
