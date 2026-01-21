package main

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func parseCommand(input string) {
	switch input {
	case "1":
		clearTerminal()
		printBookInformation()
		response := askForCommand()
		book := createBookFromInput(response)
		AddBook(book)
		clearTerminal()
		printMenu()
	case "3":
		clearTerminal()
		printEnterIsbnNumber()
		isbn := askForCommand()
		isAvailable := CheckBookAvailability(isbn)
		printIsAvailable(isAvailable)
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()
	case "4":
		clearTerminal()
		printEnterIsbnNumber()
		isbn := askForCommand()
		book := LendBook(isbn)
		printIsBookBorrowed(book)
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()
	case "5":
		clearTerminal()
		printEnterIsbnNumber()
		isbn := askForCommand()
		book := ReturnBook(isbn)
		printBookReturn(book)
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()

	case "6":
		allBooks := FindAllBooks()
		printBookList(allBooks)
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()
	case "q":
		GoodbyeMessage()
		time.Sleep(1 * time.Second)
		clearTerminal()
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
	return Book{ //Instanzieren
		ISBN:          parts[0],
		Title:         parts[1],
		Author:        parts[2],
		PublishedYear: year,
		Available:     true,
	}
}

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
