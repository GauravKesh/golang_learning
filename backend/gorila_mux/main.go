package main

import (
	"fmt"
	"net/http"
)

func serverStartLog(p string) {
	fmt.Printf("SERVER RUNNNING AT PORT %s", p)
}

func main() {
	const PORT string = "8080"

	serverStartLog(PORT) // log before starting server

	// Start server on :8080 with no handler (default mux)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
