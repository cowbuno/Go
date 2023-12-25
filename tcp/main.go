package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var requestBody RequestBody

		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, "Invalid JSON message", http.StatusBadRequest)
			return
		}

		if requestBody.Message != "" {
			fmt.Printf("Received message: %s\n", requestBody.Message)
			response := Response{Status: "success", Message: "Data successfully received"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		} else {
			http.Error(w, "Invalid JSON message", http.StatusBadRequest)
			return
		}
	})

	port := 8080
	fmt.Printf("Server is running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
