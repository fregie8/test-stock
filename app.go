package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("======= Start App =========")
	router := httprouter.New()
	router.GET("/", callHTML)
}

func callHTML(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var templateObject *template.Template
	testTemplate := templateObject.Lookup("test")
	if testTemplate != nil {
		if err := testTemplate.Execute(w, nil); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
