package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Data string `json:"message"`
}

func main() {
	fmt.Println("Iniciando o servidor Go...")

	// Configurar as opções do CORS para permitir solicitações do seu aplicativo Flutter Web
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"}, // Substitua pelo endereço do seu aplicativo Flutter Web
	})

	// Use o middleware CORS nas suas rotas
	handler := http.NewServeMux()
	handler.HandleFunc("/flutter-request", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:3000/data")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Erro na solicitação para o backend Rails", resp.StatusCode)
			return
		}

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

	http.ListenAndServe(":8080", c.Handler(handler))
}
