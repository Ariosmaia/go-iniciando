package main

import "fmt"

func main() {
	nome := "Allan"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	// Captura o que variavel digitou
	// Recebe um modificador e uma variavel
	// %d representa um inteiro
	// O que o usuario digitar quero salvar na variavel comando
	// & indica o edenreço da variavel que quero salvar
	var comando int
	//fmt.Scanf("%d", &comando)
	// Scan não precisa passar o tipo da variavel
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)
}
