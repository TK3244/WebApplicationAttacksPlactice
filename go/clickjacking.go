package main

import (
	"html/template"
	"net/http"
	"log"
)

func IndexClickJacking(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/ClickJacking/IndexClickJacking.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func NormalShop(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/ClickJacking/NormalShop.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func ClickJackingTest(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/ClickJacking/ClickJackingTest.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func PurchasedApple(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/ClickJacking/PurchasedApple.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func PurchasedOrange(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/ClickJacking/PurchasedOrange.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}
