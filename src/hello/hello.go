package main

import "fmt"
import "os"

func main() {

	exibeIntroducao()
	exibeMenu()
	leComando()

}
func exibeIntroducao() {
	nome := "Maison"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("O programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("3 - Sair")

}

func leComando() {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	switch comando {
	case 1:
		fmt.Println("Monitorando")
	case 2:
		fmt.Println("Mostrando Logs")
	case 0:
		fmt.Println("Saindo do Programa")
		os.Exit(0)
	default:
		fmt.Println("opção inválida")
		os.Exit(-1)
	}

}
