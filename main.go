package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type templateData struct {
	Service  string
	Revision string
}

var X_MASHAPE_PROXY_SECRET string
var data templateData

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	X_MASHAPE_PROXY_SECRET = os.Getenv("X_MASHAPE_PROXY_SECRET")
}

func main() {
	service := os.Getenv("K_SERVICE")
	if service == "" {
		service = "???"
	}

	revision := os.Getenv("K_REVISION")
	if revision == "" {
		revision = "???"
	}

	data = templateData{
		Service:  service,
		Revision: revision,
	}
	log.Printf("Data: %+v", data)

	mux := http.NewServeMux()
	indexHandler := http.HandlerFunc(index)
	timeHandler := http.HandlerFunc(getTime)
	mux.Handle("/", securityHandler(indexHandler))
	mux.Handle("/time", securityHandler(timeHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("The container started successfully. Listening on port %s", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func securityHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := r.Header.Get("X-Mashape-Proxy-Secret")

		//TODO: mendel remove
		// fmt.Fprintf(w, "SECRET / X: %s / %s#\n", secret, X_MASHAPE_PROXY_SECRET)

		if secret != X_MASHAPE_PROXY_SECRET {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	//TODO: mendel remove
	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "H:%s:%v\n", k, v)
	// }

	fmt.Fprintf(w, "Hello from Scout\n")
}

func getTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "time is %+v\n", time.Now().Format(time.RFC3339))
}
