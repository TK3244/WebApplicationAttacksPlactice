package main

import (
	"net/http"
	"html/template"
	"log"

	"github.com/astaxie/session"
    _ "github.com/astaxie/session/providers/memory"
)

var globalSessions *session.Manager

func IndexCSRF(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/CSRF/IndexCSRF.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func RegistrationForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/CSRF/RegistrationForm.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func LoginSession(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("username") != "" && r.FormValue("password") != "" {
		globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
		go globalSessions.GC()
		sess := globalSessions.SessionStart(w, r)
		sess.Set("username", r.FormValue("username"))
		sess.Set("password", r.FormValue("password"))
		value := map[string]string {
			"Name":     sess.Get("username").(string),
			"Password": sess.Get("password").(string),
			"ID":       sess.SessionID(),
		}
		tmpl, err := template.ParseFiles("static/CSRF/RegistrationSuccess.html")
		if err != nil {
			log.Fatalf("template.ParseFiles error err:%v", err)
		}
		err = tmpl.Execute(w, value)
		if err != nil {
			log.Fatalf("template.Execute error err:%v", err)
		}
		return
	}
	tmpl, err := template.ParseFiles("static/CSRF/RegistrationFailure.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func NewPassword(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	sess.Set("password", r.FormValue("password"))
	value := map[string]string {
		"Name":     sess.Get("username").(string),
		"Password": sess.Get("password").(string),
	}
	tmpl, err := template.ParseFiles("static/CSRF/NewPassword.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, value)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}