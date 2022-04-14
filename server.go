package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type track struct {
	TrackId      int            `json:"trackid"`
	Name         string         `json:"name"`
	AlbumId      int            `json:"albumid"`
	MediaTypeId  int            `json:"mediatypeid"`
	GenreId      int            `json:"genreid"`
	Composer     sql.NullString `json:"composer"`
	Milliseconds int            `json:"milliseconds"`
	Bytes        int            `json:"bytes"`
	UnitPrice    float64        `json:"unitprice"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []track `json:"data"`
	Message string  `json:"message"`
}

func get(w http.ResponseWriter, r *http.Request) {

	//logging request information
	fmt.Println(r)

	pathParams := mux.Vars(r)
	key := pathParams["name"]

	trackList := JsonResponse{Type: "success", Data: searchForTracks(key)}

	w.Header().Set("Content-Type", "application/json")
	//was considering using a different HTTP response code such as 204 or 404 for when there are no results from the query but doing research seemed
	//to suggest that I should still use a 200 response code when looking for a collection of results and get no result back
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trackList)

}

func searchForTracks(searchString string) []track {

	db, err := sql.Open("sqlite3", "./chinook.db")
	checkErr(err)

	defer db.Close()

	//Should use parameterized queries for security purposes
	rows, err := db.Query("SELECT * FROM Track WHERE Name like '%" + searchString + "%'")

	checkErr(err)

	defer rows.Close()

	var tracks []track

	for rows.Next() {
		var trackId int
		var name string
		var albumId int
		var mediaTypeId int
		var genreId int
		var composer sql.NullString
		var milliseconds int
		var bytes int
		var unitPrice float64
		err = rows.Scan(&trackId, &name, &albumId, &mediaTypeId, &genreId, &composer, &milliseconds, &bytes, &unitPrice)
		checkErr(err)

		tracks = append(tracks, track{TrackId: trackId, Name: name, AlbumId: albumId, MediaTypeId: mediaTypeId, GenreId: genreId, Composer: composer, Milliseconds: milliseconds, Bytes: bytes, UnitPrice: unitPrice})
	}

	return tracks
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search/{name}", get).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":4041", r))
}
