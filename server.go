package main

import (
	// "os"

	"net/http"

	// "html/template"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var h dbHandler

func init() {
	cfg := mysql.Config{
		User: "go",
		Passwd: "gopass",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "urls",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Can't create db")
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Can't establish a connection to database at dataSourceName %v", cfg.FormatDSN())
	}

	st, err := db.Prepare("select * from urls where full = ?")
	if err != nil {
		log.Fatal(err)
	}
	h = dbHandler{db, st}
}

type dbHandler struct {
	db *sql.DB
	st *sql.Stmt
}

func (h dbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if params.Has("url") {
		var full, short string
		println("url =", params.Get("url"))
		err := h.st.QueryRow(params.Get("url")).Scan(&full, &short)
		println("full", full, "\nshort", short)
		if err == sql.ErrNoRows {
			println("creating new short url")
		} else if err != nil {
			log.Fatal(err)
		}
		println("your short url:", short)
	} else {
		println("No url")
	}
	http.ServeFile(w, r, "index.html")
}

func main() {
	defer h.db.Close()
	defer h.st.Close()

	http.ListenAndServe(":8080", h)
}
