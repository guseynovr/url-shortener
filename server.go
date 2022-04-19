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
var lastID int64

func init() {
	cfg := mysql.Config{
		User:   "go",
		Passwd: "gopass",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
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
		var full, short, url string
		var id int
		url = params.Get("url")
		println("url =", url)
		err := h.st.QueryRow(url).Scan(&id, &full, &short)
		println("full", full, "short", short)
		if err == sql.ErrNoRows {
			println("creating new short url")
			short = newURL(url)
		} else if err != nil {
			log.Fatal(err)
		}
		println("your short url:", short)
	} else {
		println("No url")
	}
	http.ServeFile(w, r, "index.html")
}

func newURL(full string) string {
	// println("full", full, "lastID", lastID)
	if lastID == 0 {
		row := h.db.QueryRow("SELECT COUNT(*) FROM urls;")
		err := row.Scan(&lastID)
		if err != nil {
			panic(err)
		}
	}
	short := shorten(int(lastID + 1))
	result, err := h.db.Exec("insert into urls (Full, Short) values (?, ?)", full, short)
	if err != nil {
		//TODO: check error
		panic(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
	lastID, err = result.LastInsertId()
	if err != nil {
		//TODO: check error
		panic(err)
	}
	println("lastID new:", lastID)
	return short
}

func main() {
	defer h.db.Close()
	defer h.st.Close()

	http.ListenAndServe(":8080", h)
}
