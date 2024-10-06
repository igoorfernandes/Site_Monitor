# Site Monitor
Este é um programa escrito em Go para monitorar o status de websites. Ele verifica se os sites estão online ou offline e registra essa informação em um arquivo de log. Foi desenvolvido como parte de um curso da Alura e representa meu primeiro projeto utilizando a linguagem Go.


## Funcionalidades
- Monitoramento de múltiplos sites.
- Registra o status dos sites (online/offline) em um arquivo de log.
- Exibe os logs salvos em um arquivo.
- Usa cores no terminal para indicar sucesso (verde) ou erro (vermelho).

## Como Funciona
1. **Monitoramento**: O programa lê uma lista de sites a partir de um arquivo sites.txt. Em cada ciclo de monitoramento, ele envia uma requisição HTTP para cada site e verifica o código de status da resposta.

    - Se o status for 200 (OK), o site é considerado online.
    - Caso contrário, o site é considerado offline e o código de status é exibido.

2. **Logs**: As informações de monitoramento são registradas em um arquivo log.txt com o timestamp e o status do site.

3. **Menu**: O usuário interage com o programa através de um menu que permite iniciar o monitoramento ou exibir os logs salvos.

## Requisitos
Go instalado na sua máquina.
O pacote externo fatih/color para colorir as mensagens no terminal.
Instalação do Pacote Externo
Para instalar o pacote github.com/fatih/color, execute o seguinte comando no diretório do projeto:

```
go get github.com/fatih/color
```

## Como Usar

1. Clone o repositório ou copie o código para sua máquina.

2. Crie o arquivo sites.txt na mesma pasta do código. Liste os sites que deseja monitorar, um por linha. Exemplo:
```
https://www.google.com
https://www.github.com
https://www.stackoverflow.com
```
3. Execute o programa com o comando:
```
go run main.go
```
4. No menu, escolha uma das opções:
- 1 para iniciar o monitoramento.
- 2 para exibir os logs salvos.
- 0 para sair do programa.

## Exemplo de Uso
1. Ao escolher a opção de monitoramento, o programa testa os sites e retorna:
```
Iniciando monitoramento...
Testando site 1: https://www.google.com
Site: https://www.google.com foi carregado com sucesso!

Testando site 2: https://www.github.com
Site: https://www.github.com foi carregado com sucesso!

Testando site 3: https://www.stackoverflow.com
Site: https://www.stackoverflow.com encontrou um erro [404]
```
2. O log correspondente será salvo no arquivo log.txt:
```
05/10/2024 15:45:23 - https://www.google.com - online: true
05/10/2024 15:45:23 - https://www.github.com - online: true
05/10/2024 15:45:23 - https://www.stackoverflow.com - online: false
```
### Estrutura do Código
- main.go: O arquivo principal do programa.
- sites.txt: Lista dos sites a serem monitorados.
- log.txt: Armazena os resultados do monitoramento.
  
### Dependências
- github.com/fatih/color: Biblioteca para colorir saídas no terminal.

### Licença
O repositório está sob a licença MIT. Para mais detalhes, acesse <a href="https://github.com/igoorfernandes/Site_Monitor/blob/main/LICENSE">license</a>.
