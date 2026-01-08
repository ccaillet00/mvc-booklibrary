package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseCommand(input string) {
	switch {
	case input == "1":
		printBookInformation()
		response := askForCommand()
		book := createBookFromInput(response)
		AddBook(book)
		printMenu()
	case input == "6":
		allBooks := FindAllBooks()
		printBookList(allBooks)
		printContinue()
		askForCommand()
		printMenu()
	case input == "q":
		GoodbyeMessage()
		os.Exit(0)
	}
}

func askForCommand() string {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

func executeCommand() {
	command := askForCommand()
	parseCommand(command)
}

func createBookFromInput(response string) Book {
	clean := strings.ReplaceAll(response, ", ", ",")
	parts := strings.Split(clean, ",")
	year, _ := strconv.Atoi(parts[3])
	return Book{
		ISBN:          parts[0],
		Title:         parts[1],
		Author:        parts[2],
		PublishedYear: year,
		Available:     false,
	}
}
