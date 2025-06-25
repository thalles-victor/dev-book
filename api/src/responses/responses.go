package responses

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// JSON return a response in JSON to a request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

// Erro returns a error in JSON format
func Error(w http.ResponseWriter, statusCode int, err error) {
	if statusCode >= 500 && statusCode <= 600 {
		fmt.Println(err)

		JSON(w, statusCode, struct {
			Erro string `json:"error"`
		}{
			Erro: "Internal server error",
		})
		return
	}

	JSON(w, statusCode, struct {
		Erro string `json:"error"`
	}{
		Erro: err.Error(),
	})
}
