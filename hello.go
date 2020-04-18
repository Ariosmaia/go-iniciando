package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	//exibeIntroducao()
	exibeNome()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

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

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// Array tem tamanho fixo
	// Os arrays são limitados é melhor usar Slices
	var sites [4]string
	sites[0] = "http://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "http://www.caelum.com.br"

	fmt.Println(sites)

	site := "http://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:",
			resp.StatusCode)
	}
}

func exibeNome() {
	// Slice se adapta ai tamanho doarray
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	// Adiciona itens no slice/array
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	// cap mostra a capacidade
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
	// Toda vez que vou estourar a capacidade do slice ele dobra
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
