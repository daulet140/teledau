package main

import (
	"context"
	"cwntelegram/internal/client"
	"cwntelegram/internal/routes"
	"fmt"
	"net/http"
)

func main() {

	telegramClient := client.NewTelegramClient("6699186697:AAFntrjkj36iY-l5ZcqEBxAvf40sSiGCzpk", context.Background())

	router := routes.NewRouter(*telegramClient)

	http.HandleFunc("/webhook", router.WebhookHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
