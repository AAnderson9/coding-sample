package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type track struct {
	TrackId      int
	Name         string
	AlbumId      int
	MediaTypeId  int
	GenreId      int
	Composer     sql.NullString
	Milliseconds int
	Bytes        int
	UnitPrice    float64
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))

}

func searchForTracks(db *sql.DB, searchString string) []track {
	//Should use parameterized queries for security purposes
	rows, err := db.Query("SELECT * FROM Track WHERE Name like '%" + searchString + "%'")

	checkErr(err)

	defer rows.Close()

	tracks := make([]track, 0)

	for rows.Next() {
		track := track{}
		err = rows.Scan(&track.TrackId, &track.Name, &track.AlbumId, &track.MediaTypeId, &track.GenreId, &track.Composer, &track.Milliseconds, &track.Bytes, &track.UnitPrice)
		checkErr(err)

		tracks = append(tracks, track)
	}

	return tracks
}

func main() {
	db, err := sql.Open("sqlite3", "./chinook.db")
	checkErr(err)

	defer db.Close()

	tracks := searchForTracks(db, "and")
	fmt.Println(tracks)

	// r := mux.NewRouter()
	// r.HandleFunc("/", get).Methods(http.MethodGet)
	// log.Fatal(http.ListenAndServe(":4041", r))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
