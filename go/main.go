package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:rootpassword@tcp(db:3306)/testdb")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("resources/"))
	mux.Handle("/resources/", http.StripPrefix("/resources/", files))

	mux.HandleFunc("/", Index)

	mux.HandleFunc("/XSS/Index", IndexXSS)
	mux.HandleFunc("/XSS/XSSform", XSSform)
	mux.HandleFunc("/XSS/XSStest", XSStest)

	mux.HandleFunc("/ClickJacking/Index", IndexClickJacking)
	mux.HandleFunc("/ClickJacking/NormalShop", NormalShop)
	mux.HandleFunc("/ClickJacking/ClickJackingTest", ClickJackingTest)
	mux.HandleFunc("/ClickJacking/PurchasedApple", PurchasedApple)
	mux.HandleFunc("/ClickJacking/PurchasedOrange", PurchasedOrange)

	mux.HandleFunc("/OpenRedirect/Index", IndexOpenRedirect)
	mux.HandleFunc("/OpenRedirect/LoginForm", LoginForm)
	mux.HandleFunc("/OpenRedirect/UserLogin", UserLogin)
	mux.HandleFunc("/OpenRedirect/LoginSuccess", LoginSuccess)
	mux.HandleFunc("/OpenRedirect/TrapSite1", TrapSite1)
	mux.HandleFunc("/OpenRedirect/TrapSite2", TrapSite2)

	mux.HandleFunc("/CSRF/Index", IndexCSRF)
	mux.HandleFunc("/CSRF/RegistrationForm", RegistrationForm)
	mux.HandleFunc("/CSRF/LoginSession", LoginSession)
	mux.HandleFunc("/CSRF/NewPassword", NewPassword)

	mux.HandleFunc("/SQLinjection/Index", IndexSQLinjection)
	mux.HandleFunc("/SQLinjection/LoginForm", LoginForm2)
	mux.HandleFunc("/SQLinjection/HandleForm", HandleForm)
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/Index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles error err:%v", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("template.Execute error err:%v", err)
	}
}
