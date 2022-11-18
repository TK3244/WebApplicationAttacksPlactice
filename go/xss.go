package main

import (
	"net/http"
	"text/template"
	"log"
)

func IndexXSS(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/XSS/IndexXSS.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func XSSform(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/XSS/XSSform.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func XSStest(w http.ResponseWriter, r *http.Request) {
	value := map[string]string {
		"name": r.FormValue("name"),
	}
	tmpl, err := template.ParseFiles("static/XSS/XSStest.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, value)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}
