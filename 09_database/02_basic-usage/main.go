package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./template.gohtml"))
}

func main() {
	db, err = sql.Open("mysql", "root:1112@/myDB?charset=utf8")
	checkErr(err)
	defer db.Close()

	err = db.Ping()
	checkErr(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/search", search)
	http.HandleFunc("/delete", delete)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	checkErr(err)
}

func create(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS GoTable (name VARCHAR(20));`)
	checkErr(err)

	result, err := stmt.Exec()
	checkErr(err)

	n, err := result.RowsAffected()
	checkErr(err)

	fmt.Fprintln(res, "CREATED TABLE GoTable RECORD", n)
}

func insert(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO GoTable VALUES ("Jay")`)
	checkErr(err)

	result, err := stmt.Exec()
	checkErr(err)

	n, err := result.RowsAffected()
	checkErr(err)

	fmt.Fprintln(res, "INSERTED RECORD", n)
}

func search(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM GoTable`)
	checkErr(err)

	fmt.Fprintln(res, "RETREIVED DATA:")

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		checkErr(err)
		fmt.Fprintln(res, name)
	}
}

func delete(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM GoTable WHERE name='Jay'`)
	checkErr(err)

	result, err := stmt.Exec()
	checkErr(err)

	n, err := result.RowsAffected()
	checkErr(err)

	fmt.Fprintln(res, "DELETE RECORD", n)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
