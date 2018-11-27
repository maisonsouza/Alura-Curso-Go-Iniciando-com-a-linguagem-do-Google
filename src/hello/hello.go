package main

import "fmt"

func main() {
	nome := "Maison"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("O programa está na versão", versao)
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("3 - Sair")
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi",comando)

}
