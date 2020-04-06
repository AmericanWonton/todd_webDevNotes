package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "admin:bigPWord117@tcp(testdb1.cd8ujtto1hfj.us-east-2.rds.amazonaws.com:3306)/test_Schema?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos_test;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	/* From the documentation, Next returns the next row in the line of rows we asked for from the 'rows' variable above.
	It returns false if there's no row up next, (so basically, it's really good for loops) */
	for rows.Next() {
		/* Scan copies the columns in the current row and copies them to a destination. So we set the destination,
		(that 'name' string variable above), and point it to that */
		err = rows.Scan(&name)
		check(err)       //Check to make sure there was no error doing that above.
		s += name + "\n" //We keep adding the name returned and a newline for printing later.
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	/* According to the documentation, this prepares the table */
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()
	//These two below actually create the table.
	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	//This is similar to create
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
