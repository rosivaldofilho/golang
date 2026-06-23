package main

import (
	"fmt"
)

// String
// bool
// int | int8, int16, int64, - tamanhos diferentes de espaço de memória alocado
// uint8, uint16, uint64 - Números inteiros que não podem ser negativos
// float32, float64
// Bytes, Runes

func main() {
	var nome string = "Sujeito Programador"
	var isAdmin bool = true

	idade := 28 // := declara a variável e inicializa ela
	numero := -280
	valor := 332.55

	fmt.Println(nome)
	fmt.Println(isAdmin)
	fmt.Println(idade)
	fmt.Println(numero)
	fmt.Printf("O valor em float é: %.2f \n", valor)

	// Bytes e Runes - Devem ser usadas aspas simples

	var b byte = 'A'
	var r rune = '&'
	fmt.Println("Byte: ", b)
	fmt.Println("Rune: ", r)
}
