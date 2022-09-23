package main

func main() {
	cards := newDeckFromfile("test.txt")

	cards.shuffle()
	cards.print()
}
