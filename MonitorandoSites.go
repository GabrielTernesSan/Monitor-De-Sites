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

const Monitoramentos = 3
const delay = 5

func main() {
	BoasVindas()
	for {
		Menu()
		comandos()
	}
}

func BoasVindas() {
	fmt.Println(`Olá, este projeto tem como objetivo criar uma aplicação
de monitoramento de sites. A principal função desta aplicação é fazer uma 
verificação periódica em sites selecionados, que podem mudar de acordo 
com a necessidade, os sites que serão verificados se encontrarão em um documento txt.`)
	fmt.Println()
}

func Menu() {
	fmt.Println("1 - Iniciar Monitoranmento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Encerrar")
}

func LeComando() int {
	var comandoLido int
	fmt.Println("Digite a opção: ")
	fmt.Scanln(&comandoLido)
	return comandoLido
}

func comandos() {
	comando := LeComando()
	switch comando {
	case 1:
		IniciarMonitormento()
	case 2:
		fmt.Println("Exibindo Logs...")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Por favor, digite umas das opções!!")
		os.Exit(-1)
	}
}

func SitesArquivo() []string {

	var sites []string

	arquivo, err := os.Open(`C:\Users\Gabri\Desktop\Alura\Golang\MonitoradorSite\Sites.txt`)
	defer arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	valor := bufio.NewReader(arquivo)

	for {
		linha, err := valor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	return sites
}

func IniciarMonitormento() {
	fmt.Println("Monitorando...")

	sites := SitesArquivo()

	for j := 0; j < Monitoramentos; j++ {
		for i := range sites {
			testaSite(sites[i])
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
	fmt.Println()
}

func testaSite(sites string) {
	response, err := http.Get(sites)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", sites, "foi carregado com sucesso!")
		RegistraLog(sites, true)
	} else {
		fmt.Println("Site:", sites, "está com algum problema! Status Code:", response.StatusCode)
		RegistraLog(sites, false)
	}
}

func RegistraLog(site string, status bool) {
	arquivo, err := os.OpenFile(`C:\Users\Gabri\Desktop\Alura\Golang\MonitoradorSite\Logs.txt`, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile(`C:\Users\Gabri\Desktop\Alura\Golang\MonitoradorSite\Logs.txt`)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))
}
