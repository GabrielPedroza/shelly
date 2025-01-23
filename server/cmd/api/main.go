// This is the main entrypoint of the server. Unless you want to change the logo of Shellhacks
// in the Println, there should be no reason to touch this file.

package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/weareinit/Opal/internal/handlers"
)

// Lets keep these Println's here so we know when the server is running when running `go run main.go`
func main() {
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service on localhost:8000...")

	fmt.Println(`
  ██████  ██░ ██ ▓█████  ██▓     ██▓     ██░ ██  ▄▄▄       ▄████▄   ██ ▄█▀  ██████ 
▒██    ▒ ▓██░ ██▒▓█   ▀ ▓██▒    ▓██▒    ▓██░ ██▒▒████▄    ▒██▀ ▀█   ██▄█▒ ▒██    ▒ 
░ ▓██▄   ▒██▀▀██░▒███   ▒██░    ▒██░    ▒██▀▀██░▒██  ▀█▄  ▒▓█    ▄ ▓███▄░ ░ ▓██▄   
  ▒   ██▒░▓█ ░██ ▒▓█  ▄ ▒██░    ▒██░    ░▓█ ░██ ░██▄▄▄▄██ ▒▓▓▄ ▄██▒▓██ █▄   ▒   ██▒
▒██████▒▒░▓█▒░██▓░▒████▒░██████▒░██████▒░▓█▒░██▓ ▓█   ▓██▒▒ ▓███▀ ░▒██▒ █▄▒██████▒▒
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░░ ▒░ ░░ ▒░▓  ░░ ▒░▓  ░ ▒ ░░▒░▒ ▒▒   ▓▒█░░ ░▒ ▒  ░▒ ▒▒ ▓▒▒ ▒▓▒ ▒ ░
░ ░▒  ░ ░ ▒ ░▒░ ░ ░ ░  ░░ ░ ▒  ░░ ░ ▒  ░ ▒ ░▒░ ░  ▒   ▒▒ ░  ░  ▒   ░ ░▒ ▒░░ ░▒  ░ ░
░  ░  ░   ░  ░░ ░   ░     ░ ░     ░ ░    ░  ░░ ░  ░   ▒   ░        ░ ░░ ░ ░  ░  ░  
      ░   ░  ░  ░   ░  ░    ░  ░    ░  ░ ░  ░  ░      ░  ░░ ░      ░  ░         ░  
                                                          ░                        `)
	err := http.ListenAndServe("0.0.0.0:8000", r)
	if err != nil {
		log.Error(err)
	}
}
