package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Password string
}

func IndexSQLinjection(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/SQLinjection/IndexSQLinjection.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func LoginForm2(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/SQLinjection/LoginForm.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"ID":       r.FormValue("ID"),
		"Password": r.FormValue("password"),
	}
	req := "SELECT * FROM user WHERE id=" + value["ID"] + " AND password='" + value["Password"] + "'"
	if row := GetData(db, req); row == 0 {
		tmpl, err := template.ParseFiles("static/SQLinjection/LoginFailure.html")
		if err != nil {
			log.Fatalf("template.ParseFiles error err:%v", err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatalf("template.Execute error err:%v", err)
		}
	} else {
		tmpl, err := template.ParseFiles("static/SQLinjection/LoginSuccess.html")
		if err != nil {
			log.Fatalf("template.ParseFiles error err:%v", err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatalf("template.Execute error err:%v", err)
		}
	}

}

func GetData(db *sql.DB, req string) int {
	u := &User{}
	err := db.QueryRow(req).
		Scan(&u.ID, &u.Password)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("GetData no records.")
		return 0
	}
	if err != nil {
		log.Fatalf("GetData db.QueryRow error err:%v", err)
	}
	return 1
}
