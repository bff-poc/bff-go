package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("Iniciando o servidor Go...")

	http.HandleFunc("/flutter-request", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:3000/data")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var response Response
		if err := json.Unmarshal(body, &response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
