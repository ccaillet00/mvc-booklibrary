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
