package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"uapi/api/authorize"
	"uapi/api/authorized"
	"uapi/api/issue"
	"uapi/api/token"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func Start() {
	r := mux.NewRouter()

	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "alive")
	}).Methods(http.MethodGet)
	r.HandleFunc("/token", token.Handler).Methods(http.MethodPost)
	r.HandleFunc("/authorize", authorize.Handler).Methods(http.MethodGet)
	r.HandleFunc("/authorized", authorized.Handler).Methods(http.MethodGet)
	r.HandleFunc("/repos/{owner}/{repo}/issues", issue.Handler).Methods(http.MethodPost)

	s := &http.Server{
		Addr:         "0.0.0.0:5000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Minute,
		Handler:      r,
	}

	log.Printf("Listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
