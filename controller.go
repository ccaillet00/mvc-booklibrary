package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Controller struct {
	model Model
}

func NewController(model Model) *Controller {
	return &Controller{model: model}
}

func (c *Controller) askForCommand() string {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

func (c *Controller) parseCommand(input string) {
	switch {
	case input == "1":
		c.clearTerminal()
		printBookInformation()
		response := c.askForCommand()
		book, _ := c.createBook(response)
		c.model.AddBook(book)
		c.clearTerminal()
		printMenu()
	case input == "2":
		c.clearTerminal()
		printEnterIsbnNumber()
		isbn := c.askForCommand()
		removed := c.model.RemoveBook(isbn)
		c.clearTerminal()
		printIsBookRemoved(removed)
		printContinue()
		c.askForCommand()
		c.clearTerminal()
		printMenu()
	case input == "3":
		c.clearTerminal()
		printEnterIsbnNumber()
		isbn := c.askForCommand()
		isAvailable := c.model.CheckBookAvailability(isbn)
		c.clearTerminal()
		printIsBookAvailable(isAvailable)
		printContinue()
		c.askForCommand()
		c.clearTerminal()
		printMenu()
	case input == "4":
		c.clearTerminal()
		printEnterIsbnNumber()
		isbn := c.askForCommand()
		book := c.model.LendBook(isbn)
		printIsBookBorrowed(book)
		printContinue()
		c.askForCommand()
		c.clearTerminal()
		printMenu()
	case input == "5":
		c.clearTerminal()
		printEnterIsbnNumber()
		isbn := c.askForCommand()
		book := c.model.ReturnBook(isbn)
		printIsBookReturned(book)
		printContinue()
		c.askForCommand()
		c.clearTerminal()
		printMenu()
	case input == "6":
		c.clearTerminal()
		allBooks := c.model.FindAllBooks()
		PrintBookList(allBooks)
		printContinue()
		c.askForCommand()
		c.clearTerminal()
		printMenu()
	case input == "q":
		GoodBye()
		os.Exit(0)
	}
}

func (c *Controller) executeCommand() {
	command := c.askForCommand()
	c.parseCommand(command)
}

/*
response = 123, Matrix, Grazi, 1999
clean = 123,Matrix,Grazi,1999
parts = ["123", "Matrix", "Grazi", "1999"]
year = 1999
*/
func (c *Controller) createBook(response string) (Book, error) {
	clean := strings.ReplaceAll(response, ", ", ",")
	parts := strings.Split(clean, ",")
	if len(parts) != 4 {
		return Book{}, fmt.Errorf("erwartet 4 Felder, bekommen %d", len(parts))
	}
	year, err := strconv.Atoi(parts[3])
	if err != nil {
		return Book{}, fmt.Errorf("ungültiges Jahr %q: %w", parts[3], err)
	}

	return Book{
		ISBN:          parts[0],
		Title:         parts[1],
		Author:        parts[2],
		PublishedYear: year,
		Borrowed:      false,
	}, nil
}

func (c *Controller) clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
