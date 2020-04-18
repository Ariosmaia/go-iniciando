package main

// Pacote os fala com o sistema operacional
import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	// Nâo tem o break, ele já sai automatico
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		// 0 para sair do programa com sucesso
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		// -1 indica que tem algum problema
		os.Exit(-1)
	}
}

// Função que não retornar nada
func exibeIntroducao() {
	nome := "Allan"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

// Função com retorno
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

// O que aprendemos?

// Controle de fluxo com if
// Sua condição não fica entre parênteses e deve sempre retornar um booleano
// Controle de fluxo com switch
// Se os casos não forem atendidos, será executado o código do caso default
// Introdução às funções
// Pacote os, para encerrar o programa
