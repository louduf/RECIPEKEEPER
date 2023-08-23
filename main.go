package main

import (
	"net/http"

	"github.com/Louduff/RECIPEKEEPER/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
}
