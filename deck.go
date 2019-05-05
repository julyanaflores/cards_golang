//Funciona como a declaração de uma classe em orientação a obj
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//definindo um tipo customizado baseado em um existente (lista de string) porém com novas propriedades customizadas
//está estendendo slice string
type deck []string

//Funcao para criar um baralho novo
func newDeck() deck {
	cards := deck{}

	//Não usamos o tipo deck porq tecnicamente deck deve receber as cartas do baralho que nós estamos arrecem criando
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//Quando não há a necessidade de um index, basta usar _
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//Funcao para exibir o baralho
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Funcao para distribuir a mao de cartas
//Enquanto o primeiro par de parenteses especifica os parametros da funcao
//O segundo diz que a funcaoo deve ter dois retornos, ambos do tipo deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Funcao para converter um deck, que é um slice de strings, em uma unica string
func (d deck) toString() string {
	//Com o pacote de strings importado, usamos a funcao Join que é responsavel por transformar um slice de strings em uma unica string
	return strings.Join([]string(d), ",")

}

//Funcao para salvar no disco um arquivo.
//Será chamada por um deck e deverá passar o nome do arquivo que será salvo
//Caso a funcao de dentro retornar um erro, a funcao retornara um erro tambem
func (d deck) saveToFile(fileName string) error {

	//Apos importado o pacote ioutil para poder usar a sua funcao padrao que grava arquivos no disco,
	//São passados os parametros corretos da funcao e o ultimo é permissoes do arquivo, onde foi usado um valor padrao onde todas as permissoes sao atribuidas para qualquer um
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//Funcao para criar um novo baralho baseado em um arquivo previamente salvo no disco
//Onde o parametro é o nome do arquivo em que se quer basear
func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)

	//Estrutura de um IF na linguagem GO. Nil corresponde a nenhum valor (null)
	if err != nil {
		//Caso tenha retornado um erro, será tratado aqui
		//Podemos retornar um log do erro e um newDeck()
		//Ou podemos registrar o log de erro e fechar o programa
		fmt.Println("Error: ", err)

		//Encerrando o programa com uma funcao padrao
		//É passado por parametro 0 quando está tudo certo e qualquer numero diferente de zero quando algo deu errado
		os.Exit(1)
	}

	//Aqui é feito o caminho contrario da funcao saveToFile
	//Pois o byte é convertido em string e a string é separada em uma slice de strings
	//Que por sua vez é convertida em um deck (baralho) para retornar à funcao
	stringSlice := strings.Split(string(byteSlice), ",")
	return deck(stringSlice)

}

//Funcao para embaralhar minhas cartas
//Usando a importação de uma biblioteca chamada "math/rand" para poder usar a funcao matematica de random
func (d deck) shuffle() {
	//Criamos uma estrutura de repetição que se baseia no
	//Numero de cartas que o meu baralho passdo por parametro tem

	//É importante ressaltar que é preciso tornar aleatoria também a origem de aleatoriedade,
	//Para que sempre que for executada seja com origem em um inicio diferente, gerando resultados diferentes
	//Sendo assim, nossa variavel source do tipo source recebe um novo source (para isso usa-se a funcao UnixNano baseada no momento (time) pois é necessario um retorno do tipo int64)
	//E dessa forma nossa variavel r do tipo rand recebe um numero aleatorio baseado no numero de origem (source)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		//Obtemos uma nova posição a partir de um numero aleatorio gerado a partir
		//Do numero total de cartas que temos no baralho - 1
		newPosition := r.Intn(len(d) - 1)

		//É realizada a troca de elementos, onde a minha carta na posição i recebe a carta na nova posicao gerada
		//E a minha carta da nova posição gerada recebe a carta que estava então na posição i
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
