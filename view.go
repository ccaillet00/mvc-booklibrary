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
# q. TERMINATE BOOK LIBRARY APP`)
}

func GoodBye() {
	fmt.Println("Goodbye!")
}

func printBookInformation() {
	fmt.Println(`
Please enter all information based on this pattern:
Pattern: ISBN, Title, Author, Publishing Year:`)
}

func PrintBookList(booksToPrint []Book) {
	for i, book := range booksToPrint {
		fmt.Println(i+1, "| ISBN:", book.ISBN+",",
			"TITLE:", book.Title+",",
			"AUTHOR:", book.Author+",",
			"BORROWED:", book.Borrowed)
	}
}

func printContinue() {
	fmt.Println("\nDrücke Enter um fortzufahren...")
}

func printEnterIsbnNumber() {
	fmt.Println("Please enter the ISBN number:")
}

func printIsBookAvailable(isAvailable bool) {
	if isAvailable {
		fmt.Println("Yes, the Book is available!")
	} else {
		fmt.Println("No, the Book is not available!")
	}
}

func printIsBookBorrowed(book *Book) {
	if book == nil {
		fmt.Println("Sorry, book not available!")
		return
	}
	fmt.Println("Book:", book.ISBN, book.Title, "borrowed")
}

func printIsBookReturned(book *Book) {
	if book == nil {
		fmt.Println("Sorry, we are not expecting this book!")
		return
	}
	fmt.Println("Book:", book.ISBN, book.Title, "returned")
}

func printIsBookRemoved(removed bool) {
	if removed {
		fmt.Println("Book successfully removed!")
	} else {
		fmt.Println("Book not found!")
	}
}
