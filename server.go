package main

import (
	// "os"
	"net/http"
	// "html/template"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "go:gopass@tcp(127.0.0.1)/urls"

var (
	db sql.DB
	st sql.Stmt
)

func init() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Can't create db")
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Can't establish a connection to database at dataSourceName %v", dsn)
	}

	st, err := db.Prepare("select * from urls where short = ?")
	if err != nil {
		log.Fatal(err)
	}
	_ = st
}

func main() {
	// dsn := "go:gopass@tcp(127.0.0.1)/urls"
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal("Can't create db")
	// }
	defer db.Close()

	// if err := db.Ping(); err != nil {
	// 	log.Fatalf("Can't establish a connection to database at dataSourceName %v", dsn)
	// }
	// http.HandleFunc("/", mainPage)
	http.ListenAndServe(":8080", http.HandlerFunc(mainPage))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if params.Has("url") {
		var full, short string
		println("url =", params.Get("url"))
		err := st.QueryRow(params.Get("url")).Scan(&full, &short)
		if err != nil {
			log.Fatal(err)
		}
		println("your short url:", short)
	} else {
		println("No url")
	}
	http.ServeFile(w, r, "index.html")
}
