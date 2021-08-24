package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := "3000"
	if v := os.Getenv("PORT"); v != "" {
		port = os.Getenv("PORT")
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte("{\"message\":\"Hello Docker\"}"))
	})
	fmt.Println("Server listening on ", port)
	http.ListenAndServe(":"+port, nil)
}
