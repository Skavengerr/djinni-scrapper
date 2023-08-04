package main

import (
	"log"
	"net/http"

	configs "github.com/Skavengerrr/job-scrapper/configs"
	"github.com/Skavengerrr/job-scrapper/pkg/telegram"
)

func main() {
	cfg, err := configs.InitViper(".")
	if err != nil {
		log.Fatal(err)
	}

	botApi := telegram.InitBot(&cfg)
	bot := telegram.NewBot(botApi, &cfg)

	http.HandleFunc("https://api.telegram.org/bot6480770355:AAHBuGVvp6Rh5JsRxYgZlazYHCuGgjWEm04/setWebhook?url=https://job-scrapper-3b16a55af4da.herokuapp.com/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Update processed successfully"))
	})

	port := ":" + "3000" // Use the port number you prefer, e.g., ":3000"
	log.Fatal(http.ListenAndServe(port, nil))

	go http.ListenAndServe(port, nil)

	// Start handling incoming updates
	bot.Start()

	// Use a blocking select{} to keep the main function running
	select {}
}
