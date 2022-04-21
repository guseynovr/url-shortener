package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/* var db *sql.DB
	defer func() {
		stats := db.Stats()
		log.Println("open connections:", stats.OpenConnections)
	}() */
	dsn := "go:gopass@tcp(127.0.0.1)/urls"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Can't create db")
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Can't establish a connection to database at dataSourceName %v", dsn)
	}

	var (
		full, short string
	)
	var id int
	row := db.QueryRow("SELECT COUNT(*) FROM urls;")
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	println("id", id)
	/* st, err := db.Prepare("select * from urls where short = ?")
	if err != nil {
		log.Fatal(err)
	} */
	/* rows, err := st.Query("https://bit.ly/3JNb6D9")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&full, &short)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(full, short)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	} */
	/* stats := db.Stats()
	log.Println("open connections:", stats.OpenConnections)
	err = st.QueryRow("https://bit.ly/3JNb6D9").Scan(&full, &short)
	if err != nil {
		log.Fatal(err)
	} */
	log.Println(full, short)
}
