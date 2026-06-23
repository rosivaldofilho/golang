package main

import "fmt"

// Variáveis do tipo Var e Const

func main() {
	var nome = "Rosivaldo"
	var idade = 35
	var texto string

	fmt.Println("Variáveis")
	fmt.Println(nome)
	fmt.Println(idade)

	idade = 29
	texto = "Entendendo variáveis"

	fmt.Println("agora a idade é ", idade)
	fmt.Println(texto)
}
