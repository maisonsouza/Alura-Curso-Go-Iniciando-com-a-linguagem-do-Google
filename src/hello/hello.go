package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	for {
		exibeIntroducao()
		exibeMenu()
		leComando()
	}

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
		iniciarMonitoramento()
	case 2:
		fmt.Println("Mostrando Logs")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do Programa")
		os.Exit(0)
	default:
		fmt.Println("opção inválida")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

}

func testaSite(site string) {
	resposta, erro := http.Get(site)
	if erro != nil {
		fmt.Println("Ocorreu um erro", erro)
	}
	if resposta.StatusCode == 200 {
		fmt.Println("O site ", site, "carregou com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Aconteceu o erro foi:", resposta.StatusCode, "no site", site)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	return sites
}

func registraLog(site string, status bool) {

	arquivo, erro := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
