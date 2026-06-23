package main

import "fmt"

// Variáveis do tipo Var e Const
/*
Variáveis possuem tipagem automática (inferência de tipo)
caso seja setado o valor em uma variável que ainda não possui tipo;
Inicialização opcional
Se a variável não for inicializada,
ela receberá o valor zero do tipo de dado

*/

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
