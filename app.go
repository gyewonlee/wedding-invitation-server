package main

import (
	"database/sql"
	"net/http"

	"github.com/juhonamnam/wedding-invitation-server/httphandler"
	"github.com/juhonamnam/wedding-invitation-server/sqldb"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqldb.SetDb(db)

	mux := http.NewServeMux()
	mux.Handle("/api/guestbook", new(httphandler.GuestbookHandler))
	mux.Handle("/api/attendance", new(httphandler.AttendanceHandler))

	http.ListenAndServe(":8080", mux)
}
