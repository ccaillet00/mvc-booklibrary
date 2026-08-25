package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"mvc-booklibrary/controller"
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
	modeFlag := flag.String("mode", "", "app mode: cli, http or both (default: $APP_MODE or cli)")
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

	case "both":
		// Run the CLI and the HTTP API at the same time.
		ctrl := controller.NewController(m)
		srv := server.NewServer(m)

		go func() {
			log.Println("HTTP server listening on :8080")
			if err := http.ListenAndServe(":8080", srv.Routes()); err != nil {
				log.Fatal(err)
			}
		}()

		view.PrintMenu()
		for {
			ctrl.ExecuteCommand()
		}

	default: // "cli"
		// Run only the terminal application.
		ctrl := controller.NewController(m)
		view.PrintMenu()
		for {
			ctrl.ExecuteCommand()
		}
	}
}
