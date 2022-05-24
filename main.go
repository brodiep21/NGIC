package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func init() {
	templ = template.Must(template.ParseGlob("website/*.html"))
}

func SearchTax(w http.ResponseWriter, r *http.Request) {

	response := r.FormValue("tax")

	search, err := googlesearch.Search(context.TODO(), "tax rate of "+response)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Setting default port to %s", port)
	}
	fmt.Printf("Starting server at %s", port)

	http.HandleFunc()
	http.HandleFunc()
	http.ListenAndServe(":"+port, nil)

}

// func
