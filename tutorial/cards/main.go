package main

func main(){
	cards:=newDeck()
	cards.saveDeckToFile("hand.txt")
	
	cards=readDeckFromFile("hand.txt")
	cards.print()
	cards.shuffle()
	cards.print()
	cards.saveDeckToFile("hand.txt")
}