package main

import (
	"html/template"
	"net/http"
	"log"
)

func IndexOpenRedirect(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/OpenRedirect/IndexOpenRedirect.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func LoginForm(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		url = "http://localhost:8080/OpenRedirect/LoginSuccess"
	}
	value := map[string]string{
		"url": url,
	}
	tmpl, err := template.ParseFiles("static/OpenRedirect/LoginForm.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, value)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"id":  r.FormValue("id"),
		"pwd": r.FormValue("pwd"),
		"url": r.FormValue("url"),
	}
	if value["id"] != "" && value["pwd"] != "" {
		http.Redirect(w, r, value["url"], 302)
	}
	tmpl, err := template.ParseFiles("static/OpenRedirect/LoginFailure.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, value)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func LoginSuccess(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/OpenRedirect/LoginSuccess.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func TrapSite1(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/OpenRedirect/TrapSite1.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func TrapSite2(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"id":  r.FormValue("id"),
		"pwd": r.FormValue("pwd"),
	}
	tmpl, err := template.ParseFiles("static/OpenRedirect/TrapSite2.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, value)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}
