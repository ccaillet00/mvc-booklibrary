package view

import (
	"fmt"

	"mvc-booklibrary/model"
)

// PrintMenu prints the main menu of the library app.
func PrintMenu() {
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

// GoodBye prints a farewell message.
func GoodBye() {
	fmt.Println("Goodbye!")
}

// PrintBookInformation prompts the user for book input.
func PrintBookInformation() {
	fmt.Println(`
Please enter all information based on this pattern:
Pattern: ISBN, Title, Author, Publishing Year:`)
}

// PrintBookList prints a list of books.
func PrintBookList(booksToPrint []model.Book) {
	for i, book := range booksToPrint {
		fmt.Println(i+1, "| ISBN:", book.ISBN+",",
			"TITLE:", book.Title+",",
			"AUTHOR:", book.Author+",",
			"BORROWED:", book.Borrowed)
	}
}

// PrintContinue prompts the user to press Enter to continue.
func PrintContinue() {
	fmt.Println("\nDrücke Enter um fortzufahren...")
}

// PrintEnterIsbnNumber prompts the user for an ISBN.
func PrintEnterIsbnNumber() {
	fmt.Println("Please enter the ISBN number:")
}

// PrintIsBookAvailable prints whether a book is available.
func PrintIsBookAvailable(isAvailable bool) {
	if isAvailable {
		fmt.Println("Yes, the Book is available!")
	} else {
		fmt.Println("No, the Book is not available!")
	}
}

// PrintIsBookBorrowed prints the result of lending a book.
func PrintIsBookBorrowed(book *model.Book) {
	if book == nil {
		fmt.Println("Sorry, book not available!")
		return
	}
	fmt.Println("Book:", book.ISBN, book.Title, "borrowed")
}

// PrintIsBookReturned prints the result of returning a book.
func PrintIsBookReturned(book *model.Book) {
	if book == nil {
		fmt.Println("Sorry, we are not expecting this book!")
		return
	}
	fmt.Println("Book:", book.ISBN, book.Title, "returned")
}

// PrintIsBookRemoved prints the result of removing a book.
func PrintIsBookRemoved(removed bool) {
	if removed {
		fmt.Println("Book successfully removed!")
	} else {
		fmt.Println("Book not found!")
	}
}
