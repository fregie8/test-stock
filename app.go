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
	router.GET("/", CallHTML)
	router.GET("/ping", HandlePing)
}

func CallHTML(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var templateObject *template.Template
	testTemplate := templateObject.Lookup("test")
	if testTemplate != nil {
		if err := testTemplate.Execute(w, nil); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func HandlePing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pingStr := []byte(`{"ping":"pong"}`)
	common.NewAjax(w, r, []byte(pingStr), http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pingStr)
}
