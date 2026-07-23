# Gerador de GUIDs
Pequeno utilitário de linha de comando escrito em Go que gera GUIDs (UUIDs v4) em quantidade, direto no terminal. Surgiu de uma demanda real — precisava gerar diversos GUIDs de uma vez — e serviu como exercício para reforçar conceitos como loops, entrada/saída, tratamento de erros e uso de pacotes externos.

## Funcionalidades
Gera qualquer quantidade de GUIDs em uma única execução
Valida a entrada do usuário (aceita apenas inteiros maiores que zero)
Loop interativo: permite gerar novos lotes sem reiniciar o programa
Usa a biblioteca oficial github.com/google/uuid
Pré-requisitos

## Go
 1.16 ou superior
## Instalação
Clone o repositório e baixe as dependências:
```
git clone [github.com](https://github.com/seu-usuario/gerador-guids.git)
cd gerador-guids
go mod tidy
```

## Como usar
Execute o programa:
```
go run main.go
```

## Conceitos praticados

<ul>
    <li>Estruturas de repetição (for com condicional de saída)</li>
    <li>Entrada de dados via <code>fmt.Scanln</code></li>
    <li>Tratamento de erros e validação de input</li>
    <li>Limpeza de buffer de entrada</li>
    <li>Comparação de strings com <code>strings.EqualFold</code></li>
    <li>Integração com pacotes de terceiros (<code>github.com/google/uuid</code>)</li>
</ul>
