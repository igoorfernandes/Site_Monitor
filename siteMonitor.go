package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

const numero_de_monitoramentos = 3
const delay_entre_monitoramento = 3

func main() {

	exibeIntroducao()
	for {

		barraDeSeparacao()

		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			inicialMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("ERRO!!!")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Igor"
	versao := 1.1
	idade := 24
	//var teste float64 = 0.5

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("Olá sr.", nome, "você tem", idade, "anos")
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
}

func barraDeSeparacao() {
	colorLine := color.New(color.FgBlue).SprintFunc()
	fmt.Println(colorLine("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"))
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitormanto")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Print("Digite um número:  ")
	fmt.Scan(&comandoLido)

	return comandoLido
}

func inicialMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := lerSitesDoArquivo()

	for i := 0; i < numero_de_monitoramentos; i++ {
		fmt.Println("Teste:", i)
		time.Sleep(delay_entre_monitoramento * time.Second)
		for i, site := range sites {
			time.Sleep(3 * time.Second)
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
	}
}

func testaSite(site string) {
	colorRed := color.New(color.FgRed).SprintFunc()
	colorGreen := color.New(color.FgGreen).SprintFunc()
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(colorGreen("Site:", site, " foi carregado com sucesso!"))
		registraLog(site, true)
		fmt.Println("")
	} else {
		fmt.Println(colorRed("Site:", site, " encontrou um erro [", resp.StatusCode, "]"))
		registraLog(site, false)
		fmt.Println("")
	}

}

func lerSitesDoArquivo() []string {
	colorRed := color.New(color.FgRed).SprintFunc()
	//colorBlue := color.New(color.FgBlue).SprintFunc()
	var sites []string

	arquivos, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", colorRed(err))
	}

	leitor := bufio.NewReader(arquivos)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivos.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
