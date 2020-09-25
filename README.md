# ead_go_novatec_introducing-go

Leitura do livro *Introdução à Linguagem Go: Crie Programas Escaláveis e Confiáveis* da Novatec, escrito pelo [Caleb Doxsey](https://github.com/calebdoxsey).

Scripts de propósitos diferentes são criados em pastas separadas, sendo a nomenclatura: `capítulo_nome-sugestivo`

## Capítulos

- [x] 1. Iniciando
- [x] 2. Tipos
- [x] 3. Variáveis
- [x] 4. Estruturas de controle
- [x] 5. Arrays, fatias e mapas
- [x] 6. Funções
- [x] 7. Estruturas e interfaces
- [x] 8. Pacotes
- [x] 9. Testes
- [x] 10. Concorrência

## Run

```sh
go run folder/main.go

# ou
# cd folder
# go run .

# vim
#:so .exrc
#<leader>r

# documentação do go
godoc -http=:6060
```

## Motivos para estudar Go

- Performance
- Compilação multiplataforma
- Tipagem estática com atribuição implícita usando `:=`
- Muitas ferramentas builtin
    - Teste
    - Documentação
    - Formatação
    - Package manager

### Sites úteis

- [Go por exemplo](http://goporexemplo.golangbr.org/)
- [GoDoc](https://golang.org/doc/)
- [Your basic](https://yourbasic.org/golang/)

## Go get

Programas úteis para baixar com go:

```sh
go get golang.org/x/tools/cmd/godoc
go get golang.org/x/tools/cmd/goimports
```
