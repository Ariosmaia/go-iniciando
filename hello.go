// Indica que é o pacote principal é o sistem deve começar por ele
package main

// Pacotes de formatação
// Tem uma função TypeOf(), entre outras
import (
	"fmt"
	"reflect"
)

// Função principal do programa
// Se eu não declarar uma variavel e não usar ele dá erro
// Não preciso declarar o tipo das variaveis sem pre no GO, ele identica o tipo das variaveis
// Posso escrever as váriaveis de 3 formas:
// - var nome string = "Allan"
// - var nome = "Allan"
// - nome := "Allan"
func main() {

	// Se não passar valor para string ele mada um texto vazio
	nome := "Allan"
	// No Go se eu não atribuir valor a váriavele e deixa 0
	idade := 28
	// Se não passar nada ele manda 0.0
	versao := 1.1
	// Primeira letra da função é maisuclo
	// , para concatenar
	fmt.Println("Olá sr., Allan Rios", nome, "Sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
}

// Para compilar:
// go build hello.go
// ./hello

// Para facilitar e não ficar dando o comando duas vezes:
// Compila e executa
// go run hello.go

// ; é opcional em Go
// Go ele já tem boas práticas de inicio

// O que aprendemos?
// Instalação do Go
// Go Workspace
// A pasta bin para os arquivos binários, compilados
// A pasta src para o código fonte
// A pasta pkg para os pacotes compartilhados entre as aplicações
// Instalação da extensão do Go no Visual Studio Code
// Com isso temos autocomplete, detecção de erros, etc
// Convenções da linguagem
// Implementação do Hello World
// Compilando e executando um programa em Go
