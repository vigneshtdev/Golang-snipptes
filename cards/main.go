package main

func main() {
	//var card string = "Ace of Spades"
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	//cards := 	[]string{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	/*
		for i, card := range cards {
			fmt.Println(i, card)
		}
	*/
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
