package main

func main() {
	repo := NewJSONRepository("books.json")
	model := NewModel(repo)
	controller := NewController(model)

	printMenu()
	for {
		controller.executeCommand()
	}
}
