package main

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

var Tax struct {
	Rate string
}

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("*.html"))
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.html", nil)
}

func SearchTax(w http.ResponseWriter, r *http.Request) (string, error) {

	response := r.FormValue("tax")
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return "", errors.New("Could not respond to form method " + r.Method)
	}

	search, err := googlesearch.Search(context.TODO(), "tax rate of "+response)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range search {
		if strings.Contains(v.Description, "%") {
			pos := strings.Index(v.Description, "%")
			if pos == -1 {
				fmt.Println("Could not locate tax")
			}
			Tax.Rate = v.Description[(pos - 5):pos]
			break
		}

	}
	return Tax.Rate, nil
}

func main() {

	ch := make(chan string)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Setting default port to %s", port)
	}

	http.HandleFunc("/", LandingPage)

	defer close(ch)
	ch <- Tax.Rate

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":"+port, nil)
	fmt.Printf("Starting server at %s", port)

}

// func
