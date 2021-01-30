package app

import (
	"log"
	"net/http"
)

// App to fill out
type App struct {
	Verbose bool
	Addr    string
}

// Run app
func (a *App) Run() error {
	if a.Verbose {
		log.Println("Verbose logging")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(a.Addr, nil)
	return nil
}
