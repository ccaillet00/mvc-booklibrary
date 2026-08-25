package controller

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"mvc-booklibrary/model"
	"mvc-booklibrary/view"
)

// Controller handles user input and coordinates the model and view.
type Controller struct {
	model *model.Model
}

// NewController creates a new Controller with the given model.
func NewController(m *model.Model) *Controller {
	return &Controller{model: m}
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
		view.PrintBookInformation()
		response := c.askForCommand()
		book, _ := c.CreateBook(response)
		c.model.AddBook(book)
		c.clearTerminal()
		view.PrintMenu()
	case input == "2":
		c.clearTerminal()
		view.PrintEnterIsbnNumber()
		isbn := c.askForCommand()
		removed := c.model.RemoveBook(isbn)
		c.clearTerminal()
		view.PrintIsBookRemoved(removed)
		view.PrintContinue()
		c.askForCommand()
		c.clearTerminal()
		view.PrintMenu()
	case input == "3":
		c.clearTerminal()
		view.PrintEnterIsbnNumber()
		isbn := c.askForCommand()
		isAvailable := c.model.CheckBookAvailability(isbn)
		c.clearTerminal()
		view.PrintIsBookAvailable(isAvailable)
		view.PrintContinue()
		c.askForCommand()
		c.clearTerminal()
		view.PrintMenu()
	case input == "4":
		c.clearTerminal()
		view.PrintEnterIsbnNumber()
		isbn := c.askForCommand()
		book := c.model.LendBook(isbn)
		view.PrintIsBookBorrowed(book)
		view.PrintContinue()
		c.askForCommand()
		c.clearTerminal()
		view.PrintMenu()
	case input == "5":
		c.clearTerminal()
		view.PrintEnterIsbnNumber()
		isbn := c.askForCommand()
		book := c.model.ReturnBook(isbn)
		view.PrintIsBookReturned(book)
		view.PrintContinue()
		c.askForCommand()
		c.clearTerminal()
		view.PrintMenu()
	case input == "6":
		c.clearTerminal()
		allBooks := c.model.FindAllBooks()
		view.PrintBookList(allBooks)
		view.PrintContinue()
		c.askForCommand()
		c.clearTerminal()
		view.PrintMenu()
	case input == "q":
		view.GoodBye()
		os.Exit(0)
	}
}

// ExecuteCommand reads a command from the user and processes it.
func (c *Controller) ExecuteCommand() {
	command := c.askForCommand()
	c.parseCommand(command)
}

/*
response = 123, Matrix, Grazi, 1999
clean = 123,Matrix,Grazi,1999
parts = ["123", "Matrix", "Grazi", "1999"]
year = 1999
*/

// CreateBook parses a comma-separated string into a Book.
func (c *Controller) CreateBook(response string) (model.Book, error) {
	clean := strings.ReplaceAll(response, ", ", ",")
	parts := strings.Split(clean, ",")
	if len(parts) != 4 {
		return model.Book{}, fmt.Errorf("erwartet 4 Felder, bekommen %d", len(parts))
	}
	year, err := strconv.Atoi(parts[3])
	if err != nil {
		return model.Book{}, fmt.Errorf("ungültiges Jahr %q: %w", parts[3], err)
	}

	return model.Book{
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
