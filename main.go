package main

import (
	"fmt"
)

type Book struct {
	Title         string
	Author        string
	ISBN          string
	PublishedYear int
	Available     bool
}

var library []Book

/*
func askUser(question string) string {
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}*/

func main() {

	library = append(library, Book{
		Title:         "The Go Programming Language",
		Author:        "Alan A. A. Donovan",
		ISBN:          "978-0134190440",
		PublishedYear: 2015,
		Available:     true,
	})

	fmt.Println(library[0])
	/*
		fmt.Println("=== Welcome to our library ===")

		for {
			fmt.Println("\n--- HAUPTMENÜ ---")
			fmt.Println("1. ADD A BOOK TO LIBRARY")
			fmt.Println("2. REMOVE A BOOK FROM LIBRARY")
			fmt.Println("3. CHECK AVAILABLE BOOKS")
			fmt.Println("4. LEND A BOOK")
			fmt.Println("5. RETURN A BOOK")
			fmt.Println("6. VIEW ALL BOOKS")
			fmt.Println("7. EXIT")

			auswahl := askUser("Deine Wahl (1-7):")

			switch auswahl {
			case "1":
				getBookInfo()
			default:
				fmt.Println("❌ Ungültige Auswahl! Bitte wähle 1-6.")
			}
		}*/
}
