package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

type Tax struct {
	Rate string
}

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("*.html"))
}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.html", nil)
}

var t Tax

func SearchTax(w http.ResponseWriter, r *http.Request) {

	response := r.FormValue("tax")

	fmt.Println(response)
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
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
			t.Rate = v.Description[(pos - 5):pos]
			break
		}

	}
	templ.ExecuteTemplate(w, "index.html", t.Rate)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
		log.Printf("Setting default port to %s", port)
	}

	http.HandleFunc("/", LandingPage)
	http.HandleFunc("?", SearchTax)

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":"+port, nil)
	fmt.Printf("Starting server at %s\n", port)

}

// func
