package main

func main() {

	//Especificando um slice em go
	cards := newDeck()

	//Atribui os retornos da funcao para cada variavel criada respectivamente
	//Ambas as variaveis serao do tipo deck porq estao recebendo esse tipo de valor
	//hand, remainingCards := deal(cards, 5)

	//Imprindo os valores da lista
	//Funciona apenas porque cards é uma lista do tipo deck e toda variavel desse tipo ganhou acesso ao metodo print()
	//cards.print()

	//hand.print()
	//remainingCards.print()

	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting)) //converte um string em um array de bytes que representam a mesma cadeia de caracteres

	//Cria uma mao de cartas a partir de um arquivo existente
	//cards := newDeckFromFile("my_cards")
	//cards.print()

	//Após obter meu baralho, vou embaralhar ele e mostrá-lo:
	cards.shuffle()
	cards.print()
}
