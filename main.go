package main // pacote principal da app

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var name string // variavel criada a nivel de packge, fora da func, ficam disponivel pra qualquer func dentro do packge

// tipos de declaracao de variaveis

func types() {
	var name string // declaracao de variavel com var
	name = "aninha"
	fmt.Println("Hello,", name)

	var name2 = "aninha" // declaracao de variavel com var e atribuicao de valor
	fmt.Println("Hello,", name2)

	name3 := "aninha" // declaracao de variavel com inferencia de tipo
	fmt.Println("Hello,", name3)

	var total int
	total++
	fmt.Println("total:", total)
}

// Funções
// Funções são blocos de código que podem ser chamados e executados em qualquer parte do código.
// é possivel retornarmos mais de um valor em uma função
// todas essas funções criadas só possuem escopo do package, ou seja, só podem ser chamadas dentro do package.
// para utilizalas em outros packages, é necessário exporta-las, ou seja, a primeira letra da função tem que ser maiuscula
// No go nao existe public, private, protected, tudo é publico por padrão, a menos que seja exportado com a primeira letra minuscula

func functions(name2 string) {
	fmt.Println("Hello,", name2)
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) int {
	i, _ := strconv.Atoi(b)

	return a + i
}

// fluxos de condicionais
// fluxos que condicionam algo.
// if, switch case

func conditionals() {
	a, b := 10, 13

	if a > b {
		fmt.Println("a é maior que b")
	} else if a < b {
		fmt.Println("b é maior que a")
	}
}

// outra forma de validarmos é criar variaveis dentro do escopo do if. em alguns casos isso é muito bom pois voce economiza memoria, otimiza o codigo e deixa mais legivel.

func openfile() {
	file, error := os.Open("hello")
	if error != nil {
		log.Panic(error)
	}

	data := make([]byte, 100) // array de bytes com 100 posições. por que um array de byte? pq a função que vamos chamar espera um array de bytes do arquivo que vamos abrir.
	if _, error := file.Read(data); error != nil {
		log.Panic(error)
	}
	log.Println("data:", string(data))
}

// caso de uso com switch case:

// DICA: quando notar que está tendo o uso várias vezes de if else, é um bom sinal que você pode usar o switch case!

var (
	cara, coroa int
)

func flipcoin(side string) {
	switch side {
	case "cara":
		cara++
	case "coroa":
		coroa++

	default:
		fmt.Println("não foi possivel identificar o lado da moeda, deve ter caido em pé!")
	}
}

// looping - Iterar várias vezes sobre o mesmo bloco de código. Veremos as principais formas
// for

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}
}

// Ex com um Array de nomes

func loop_ex2() {
	names := []string{"aninha", "joao", "maria"}
	for i := 0; i < len(names); i++ {
		fmt.Println("i:", names[i])
	}
}

// segunda forma de fazer um loop, usando range que é uma função que retorna o indice e o valor do array

func loop_ex3() {
	names := []string{"aninha", "Raissa", "maria"}
	for _, name := range names {
		fmt.Println(name)
	}
}

// terceira forma, como se fosse um while - então temos que criar uma variavel e incrementar ela dentro do bloco de código

func loop_ex4() {
	names := []string{"aninha", "Raissa", "maria"}
	var i int
	for i < len(names) {
		fmt.Println(names[i])
		i++
	}
}

// função principal da app, temos que chamar as outras funções dentro dela para que a app funcione
func main() {
	//name = "aninha"
	//fmt.Println("Hello,", name)
	//sum(1, 2)
	//fmt.Println("Soma:", sum(1, 2))
	//fmt.Println("total:", convertAndSum(90, "2"))
	//conditionals()
	//openfile()
	//flipcoin("")
	//loop()
	//loop_ex2()
	//loop_ex3()
	loop_ex4()
}
