package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	grace "gopkg.in/paytm/grace.v1"
)

func main() {
	log.Println("======= Start App =========")
	router := httprouter.New()
	router.GET("/", CallHTML)
	router.GET("/ping", HandlePing)

	n := negroni.New()
	n.UseHandler(router)

	log.Fatal(grace.Serve(fmt.Sprintf(":%s", "9000"), n))

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pingStr)
}
