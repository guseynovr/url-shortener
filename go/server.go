package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path"

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
	response := formResponse(params)
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, response)
	if err != nil {
		log.Println("template.Execute:", err)
	}
	// http.ServeFile(w, r, "index.html")
}

func formResponse(params url.Values) string {
	if params.Has("url") {
		var full, short, rawURL string
		var id int
		rawURL = params.Get("url")
		u, err := validateURL(rawURL)
		if err != nil {
			return "Invalid URL"
		}
		err = h.st.QueryRow(u).Scan(&id, &full, &short)
		// println("full", full, "short", short)
		if err == sql.ErrNoRows {
			log.Println("creating new short url")
			short = newURL(u)
		} else if err != nil {
			log.Fatal(err)
		}
		log.Println("your short url:", short)
		return "Your short url: 127.0.0.1:8080/r/" + short
	} else {
		log.Println("No url")
		return "Your short url will appear here"
	}
}

func validateURL(rawURL string) (string, error) {
	tempU, err := url.Parse(rawURL)
	if err != nil {
		log.Println(err)
	}
	if tempU.Scheme == "" {
		tempU.Scheme = "http"
	}
	u, err := url.ParseRequestURI(tempU.String())
	if err != nil {
		log.Println(err)
		return "", err
	}
	return u.String(), nil
}

func newURL(full string) string {
	// println("full", full, "lastID", lastID)
	if lastID == 0 {
		row := h.db.QueryRow("SELECT COUNT(*) FROM urls;")
		err := row.Scan(&lastID)
		if err != nil {
			//TODO: check error
			log.Fatal(err)
		}
	}
	short := shorten(int(lastID + 1))
	result, err := h.db.Exec("insert into urls (Full, Short) values (?, ?)", full, short)
	if err != nil {
		//TODO: check error
		log.Fatal(err)
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
		log.Fatal(err)
	}
	println("lastID new:", lastID)
	return short
}

func redirect(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	short := path.Base(r.URL.Path)
	println("1short:", short)
	id := resolve(short)
	println("short:", short, "id:", id)
	row := h.db.QueryRow("select * from urls where id = ?", id)

	var full string
	err := row.Scan(&id, &full, &short)
	if err == sql.ErrNoRows {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Fatal(err)
	}
	println(full)
	http.Redirect(w, r, full, http.StatusMovedPermanently)
}

func main() {
	defer h.db.Close()
	defer h.st.Close()

	http.HandleFunc("/r/", http.HandlerFunc(redirect))
	http.HandleFunc("/", http.HandlerFunc(h.ServeHTTP))
	http.ListenAndServe(":8080", nil)
}
