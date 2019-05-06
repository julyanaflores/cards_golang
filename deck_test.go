package main

import (
	"os"
	"testing"
)

//Como sabemos o que precisamos testar?
//É importante testar todas funções do código, assim como tudo o que contém lógica

//ATENÇÃO: as funções sempre começam com letra maiuscula aqui

//Para testa uma função, por exemplo a newDeck()
//Será necessário criar uma função de teste dentro desse arquivo
//E nessa função escrever códigos que garantem os resultados que a função deveria retornar

//Testando a função newDeck()
func TestNewDeck(t *testing.T) {
	d := newDeck()

	//1° teste: garantir que o meu baralho é criado com o numero correto de cartas (16 no caso)
	//O parametro t do tipo teste é responsável por dizer que algo de errado não está certo
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//2° teste: garantir que a primeira carta do baralho seja Ace of Spades (no caso)
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spaces, but got %v", d[0])
	}

	//3° teste: garantir que a ultima carta do baralho seja Four of Clubs (no caso)
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubes, but gor %v", d[len(d)-1])
	}
}

//Testando as funções saveToFile() e newDeckFromFile()
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//Começamos essa parte dos testes garantindo que não haja nenhum arquivo que possa conflitar com o que estamos procurando
	//Então removemos caso exista algum
	os.Remove("_deckdesting")

	deck := newDeck()
	deck.saveToFile("_deckdesting")
	loadedDeck := newDeckFromFile("_deckdesting")

	//1° teste: garante que o baralho carregado do arquivo possui o numero correto de cartas
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_deckdesting")
}
