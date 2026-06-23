package main

import "fmt"

/*
Variáveis do tipo Var e Const

Variáveis
Possuem tipagem automática (inferência de tipo)
Caso seja setado o valor em uma variável que ainda não possui tipo;
Inicialização opcional
Se a variável não for inicializada,
ela receberá o valor zero do tipo de dado

Const (Constante)
Não pode ser alterada após ser inicializada
Também possui tipagem automática
Precisa ser inicializada obrigatóriamente
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

	//CONST

	const pi = 3.14159
	fmt.Print(pi)
}
