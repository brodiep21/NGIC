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

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("*.html"))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Setting default port to %s", port)
	}

	http.HandleFunc("/", theOneHandler)
	// http.HandleFunc("/help", theOneHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	http.ListenAndServe(":"+port, nil)
}

func theOneHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.html", nil)
}

type Data struct {
	Tax string
}

func post(w http.ResponseWriter, r *http.Request) {
	zipCode := r.FormValue("tax")
	tax, err := searchForTaxRate(zipCode)
	if err != nil {
		fmt.Println(err)
	}

	data := Data{
		Tax: tax,
	}
	templ.ExecuteTemplate(w, "index.html", data)
}

func searchForTaxRate(zipCode string) (string, error) {

	search, err := googlesearch.Search(context.TODO(), "tax rate of "+zipCode)
	if err != nil {
		return "", err
	}

	for _, v := range search {
		if strings.Contains(v.Description, "%") {
			pos := strings.Index(v.Description, "%")
			if pos == -1 {
				break
			}
			tax := v.Description[(pos - 6) : pos+1]
			tax = strings.TrimPrefix(tax, "is ")
			tax = strings.TrimPrefix(tax, " is ")
			tax = strings.TrimPrefix(tax, "is")
			tax = strings.TrimPrefix(tax, " ")
			tax = strings.TrimPrefix(tax, "s")
			tax = strings.TrimPrefix(tax, "a")
			return tax, nil
		}
	}

	return "", errors.New("could not locate tax")
}
