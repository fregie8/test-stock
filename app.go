package main

import (
	"html/template"
	"log"
)

func main() {
	log.Println("======= Start App =========")
	var template *template.Template
	template.Lookup("test")
}
