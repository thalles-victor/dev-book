package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger write information in the terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Verificar se o usuário fazendo a requisição está autenticado
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validando...")
		next(w, r)
	}
}
