package main

import "fmt"

func printMenu() {
	fmt.Println(`
###########################################
#******* WELCOME TO OUR LIBRARY ***********
# 1. ADD A BOOK TO LIBRARY
# 2. REMOVE A BOOK FROM LIBRARY
# 3. CHECK AVAILABILITY
# 4. LEND A BOOK
# 5. RETURN A BOOK
# 6. VIEW ALL BOOKS
#
# q. TERMINATE BOOK LIBRARY APP
`)
}

func GoodbyeMessage() {
	fmt.Println("Exiting the Book Library App. Goodbye!")
}

func printBookInformation() {
	fmt.Println(`Pleae enter all information based on this patern:
Pattern: ISBN, Title, Author, PublishedYear:`)
}

func printBookList(bookToPrint []Book) {
	for i, book := range bookToPrint {
		fmt.Println(i+1, "| ISBN:", book.ISBN+",",
			"Title", book.Title+",",
			"Author", book.Author+",",
			"Available", book.Available)
	}
}

func printContinue() {
	fmt.Println("\nDrücke Enter um fortzufahren...")
}

func printEnterIsbnNumber() {
	fmt.Println("Please enter the ISBN number:")
}

func printIsAvailable(isAvailable bool) {
	if isAvailable {
		fmt.Println("Yes book is available!")
	} else {
		fmt.Println("No the book is not available!")
	}
}

func printIsBookBorrowed(book *Book) {
	if book == nil {
		fmt.Println("Sorry, book not availale!")
		return
	}
	fmt.Println("Book:", book.ISBN, book.Title, "borrowed")
}

func printBookReturn(book *Book) {
	if book == nil {
		fmt.Println("Sorry, cant return book")
		return
	}
	fmt.Println("Book", book.ISBN, book.Title, "returned")
}

func printIsBookRemoved(removed bool) {
	if removed {
		fmt.Println("Book successfully removed!")
	} else {
		fmt.Println("Book not found!")
	}
}
