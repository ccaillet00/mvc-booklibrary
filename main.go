package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"mvc-booklibrary/controller"
	"mvc-booklibrary/gin_server"
	"mvc-booklibrary/model"
	"mvc-booklibrary/repository"
	"mvc-booklibrary/server"
	"mvc-booklibrary/view"
)

// resolveMode determines the app mode:
// 1. -mode CLI flag (highest priority)
// 2. APP_MODE environment variable (useful for Docker)
// 3. default: "cli"
// Valid modes: "cli", "http", "both"
func resolveMode(flagValue string) string {
	if flagValue != "" {
		return strings.ToLower(flagValue)
	}
	if env := os.Getenv("APP_MODE"); env != "" {
		return strings.ToLower(env)
	}
	return "cli"
}

func main() {
	modeFlag := flag.String("mode", "", "app mode: cli, http or gin-http (default: $APP_MODE or cli)")
	flag.Parse()
	mode := resolveMode(*modeFlag)

	repo := repository.NewJSONRepository("books.json")
	m := model.NewModel(repo)

	switch mode {
	case "http":
		// Run only the HTTP API (blocking).
		srv := server.NewServer(m)
		log.Println("HTTP server listening on :8080")
		log.Fatal(http.ListenAndServe(":8080", srv.Routes()))

	case "cli":
		// Run only the terminal application.
		ctrl := controller.NewController(m)
		view.PrintMenu()
		for {
			ctrl.ExecuteCommand()
		}

	case "gin-http":
		// Run the HTTP API using the Gin framework (blocking).
		srv := gin_server.NewGinServer(m)
		log.Println("HTTP server listening on :8080")
		log.Fatal(http.ListenAndServe(":8080", srv.Routes()))
	}
}
