package main

func main() {
	addBookForTesting()

	printMenu()
	for {
		executeCommand()
	}
}

func addBookForTesting() {
	AddBook(Book{
		ISBN:          "978-3-16-148410-0",
		Title:         "The Go Programming Language",
		Author:        "Cédric",
		PublishedYear: 2015,
		Available:     false,
	})
	{
		AddBook(Book{
			ISBN:          "978-3-16-148410-1",
			Title:         "Bingo II",
			Author:        "Bongo",
			PublishedYear: 2004,
			Available:     false,
		})
	}
}
