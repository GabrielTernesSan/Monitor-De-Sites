# Monitor de Sites

​	O objetivo principal deste projeto é aplicar os conhecimentos adquiridos sobre a linguagem até o momento, como:

- Funções
- Slices
- Leitura, escrita e manipulação de documentos
- Loopins
- Tratamento de erros
- Bibliotecas
- Requisições HTTP

​	Bom, então vamos lá, a principal função desta aplicação é fazer uma verificação periódica em sites selecionados, que podem mudar de acordo com a necessidade do usuário, os sites que serão verificados se encontrarão em um documento txt.

![Sites](https://github.com/GabrielTernesSan/Projetos/blob/master/Imagens/Sites.png)	

​	Quando o usuário rodar a aplicação uma mensagem de boas vindas com uma breve explicação de sua função será exibida além de um menu de controle, onde apresentará 3 opções:

1. Iniciar Monitoramento
2. Exibir os Logs
3. Encerrar Programa

![Menu](https://github.com/GabrielTernesSan/Projetos/blob/master/Imagens/MenuMonitorando.png)

​	O programa foi dividido em algumas funções simples de forma que a leitura e a reprodução sejam feitas de forma simples e organizada. A aplicação foi dividida em 9 funções pequenas e objetivas.

1. BoasVindas()
2. Menu()
3. LeComando()
4. comandos()
5. IniciarMonitoramento()
6. testaSite()
7. SitesAquivo()
8. RegistraLog()
9. imprimeLogs()

## Introdução

   ``````go
   func BoasVindas() {
   	fmt.Println(`Olá, este projeto tem como objetivo criar uma aplicação
   de monitoramento de sites. A principal função desta aplicação é fazer uma 
   verificação periódica em sites selecionados, que podem mudar de acordo 
   com a necessidade, os sites que serão verificados se encontrarão em um documento txt.`)
   	fmt.Println()
   }
   ``````

   A função BoasVinda() é uma função sem retorno que basicamente tem como função exibir uma mensagem explicativa sobre como a aplicação funciona.

   ## Menu

   ``````go
   func Menu() {
   	fmt.Println("1 - Iniciar Monitoranmento")
   	fmt.Println("2 - Exibir Logs")
   	fmt.Println("0 - Encerrar")
   }
   ``````

   A função Menu() é uma função sem retorno que simplesmente "printa" o menu.

   ## Ler os comandos

   ``````go
   func LeComando() int {
   	var comandoLido int
   	fmt.Println("Digite a opção: ")
   	fmt.Scanln(&comandoLido)
   	return comandoLido
   }
   ``````

   A função LeComando() é uma função que retorna um inteiro representando uma das opções presentes no menu, após o usuário digitar um número ela é salva na variável "comandoLido" que é retornada.

   ## Comandos

   ``````go
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
   ``````

   A função comandos() é a principal função do programa e é por ele que o usuário irá "navegar". 

   Ao executar a função comandos() a função LeComando() é chamada e atribuida à variável "comando" que receberá o valor digitado pelo usuário. Em seguida a variável entra em uma estrutura de controle, onde é comparada com os "cases", o case 1 chama a função IniciarMonitoramento(), o case 2 chamará a função imprimeLogs() ambas funções serão explicadas posteriormente. Já o case 0 tem a função sair do programa utilizando uma função do package "os", a função os.Exit(), sair faz com que o programa atual seja encerrado com o código de status fornecido. Convencionalmente, o código zero indica sucesso, diferente de zero, um erro.

   Caso nenhuma das opções não seja selecionada o programa é encerrado com uma mensagem de erro.

## Sites do arquivo

````go
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
````

​	A função SitesArquivo() é responsável por fazer a leitura do documento contendo os sites a serem monitorados e em seguida atribuir a uma variável do tipo slice.

​	Para fazer isso temos que criar uma variável do tipo slice de string que irá receber estes sites mais adiante. Em seguida criamos uma variável responsável por abrir o arquivo contendo as URLs, a variável "arquivo", para abrir o arquivo utilizamos a função `os.Open()` do package "os". A função `os.Open()`, se for bem-sucedida, os métodos dos arquivo retornado podem ser usados para leitura e se houver um erro irá devolver uma mensagem. Se houver um erro o exibimos na tela.

​	Agora que abrimos o arquivo, vamos lê-lo e para isso criamos uma variável "valor" e utilizaremos o package "bufio", o package bufio facilita a leitura de arquivos, a função `bufio.NewReader` retorna um leitor, este leitor vai nos ajudar a usar algumas ferramentas para "passear" no nosso arquivo, por exemplo lermos nosso arquivo linha a linha. 

​	Para isso vamos criar uma estrutura de repetição `for` e usamos a função `ReadString()` que lê a linha do arquivo e retorna uma string, com ela também podemos definir um byte delimitador, neste nosso caso é uma quebra de linha ('\n'). A função `ReadString()` também retorna um erro, um erro específico, o EOF (End Of File) que indica que o arquivo chegou ao fim, utilizaremos este erro para sair da estrutura de repetição usando o `break`.

​	A função `strings.Trimspace` retira qualquer espaço e/ou quebras de linha desnecessários da nossa string.

​	Ao final deste `for` atribuímos as strings à slice "sites", que receberá todos os sites informados no documento.

## Monitoramento

`````go
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
`````

​	A função IniciarMonitoramento() começa chamando a função SitesArquivo() e atribui à variável "sites", em seguida o programa entra em dois loopings, o primeiro looping indica quantas vezes a verificação dos sites serão feitas, esse controle é feito com a constante "Monitoramento". 

​	Já o segundo looping faz a verificação de todos os sites um a um presentes no documento chamando a função testaSite(). Após sair do segundo "for" um tempo é dado para uma nova verificação, esse delay é dado utilizando o package time e a função ```time.Sleep```, utilizando uma segunda constante, a "delay", juntamente com a função ```time.Second``` interrompe o loop.

## Testando os sites

``````go
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
``````

​	Depois de ler o arquivo do usuário com os links precisamos verificá-los se eles estão online, para fazer isso utilizaremos um package do Go responsável pela parte web, o package "net/http". Para fazer o teste criamos uma função "testaSite" que irá receber a lista de sites, para testar os sites vamos usar a função `http.Get(sites)`, ela irá fazer uma requisição http, o retorno da função será o link da página carregada,

```go
&{200 OK 200 HTTP/2.0 2 0 map[X-Cloud-Trace-Context:[6f3fa7e590ac68bd43d76c82a67df476] Date:[Tue, 13 Jun 2017 21:20:36
GMT] Server:[Google Frontend] X-Ua-Compatible:[IE=edge,chrome=1] Expires:[Tue, 13 Jun 2017 21:50:36 GMT] Content-Type:[
text/html] Cache-Control:[public, max-age=1800] Age:[1298] X-Dns-Prefetch-Control:[on]] 0xc4200d4900 -1 [] false true m
ap[] 0xc42000a800 0xc4203a3080}
```

 mas para nós a única coisa que realmente nos interessa é o status code da requisição http. Então vamos usar uma outra função do package "net/http", a função `StatusCode`, se a resposta da página for 200 o que significa que ela carregou corretamente, fazemos o registro do site e indicamos que ela carregou com sucesso passando como argumento "true" juntamente com o site verificado. Caso haja algum erro fazemos o registro indicando o site e qual foi o status code devolvido pela página, passamos como argumento o site e o argumento "false".

## Registrando os Logs

````go
func RegistraLog(site string, status bool) {
	arquivo, err := os.OpenFile(`C:\Users\Gabri\Desktop\Alura\Golang\MonitoradorSite\Logs.txt`, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer arquivo.Close()

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
}
````

​	A função RegistraLogs() ficará responsável por salvar em um arquivo txt os testes realizados na função "testaSite()", para fazer isso abrimos o arquivo onde esses registros serão feitos. Usaremos novamente o package "os" caso a função `OpenFile` identificar que o arquivo não existe ele o criará isso é possível através das Flags da função `OpenFile` .

`````go
const (
	O_RDONLY int = syscall.O_RDONLY // abra o arquivo somente leitura.
	O_WRONLY int = syscall.O_WRONLY // abra o arquivo somente para escrita.
	O_RDWR   int = syscall.O_RDWR   // abra o arquivo para ler e escrever.
	// Os valores restantes podem ser inseridos para controlar o comportamento.
	O_APPEND int = syscall.O_APPEND // anexar dados ao arquivo ao escrever.
	O_CREATE int = syscall.O_CREAT  // crie um novo arquivo se não houver nenhum.
	O_EXCL   int = syscall.O_EXCL   // usado com O_CREATE, o arquivo não deve existir.
	O_SYNC   int = syscall.O_SYNC   // aberto para E / S síncrona.
	O_TRUNC  int = syscall.O_TRUNC  // truncar arquivo normal gravável quando aberto
)
`````

​	Usaremos 3 destas Flags a `os.RDWR` para podermos ler e escrever no arquivo em questão, a `os.O_CREATE` para criarmos o arquivo caso não exista e a flag `os.O_APPEND` para editar o arquivo. E o último parâmetro desta função são as permissões, a que vamos usar é a 0666 uma permissão padrão.

## Imprimindo os Logs

````go
func imprimeLogs() {
	arquivo, err := ioutil.ReadFile(`C:\Users\Gabri\Desktop\Alura\Golang\MonitoradorSite\Logs.txt`)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))
}
````

​	A função `imprimeLogs()` simplesmente exibe os logs registrados. Usaremos o package io/ioutil e a função `ioutil.ReadFile` para abrir o arquivo e atribuir á variável "arquivo", a variável irá receber um array de bytes que convertemos para string dentro da função `fmt.Println()`
