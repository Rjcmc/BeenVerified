package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"goji.io"
	"goji.io/pat"
	
)

type Song struct {
	Song    string  "json:song"
	Artist  string  "json:artist"
	Genre   string  "json:genre"
	Lenght  int     "json:lenght"
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/artist/:artistName"), searchArtist)
	mux.Use(logging)
	http.ListenAndServe("localhost:8000", mux)
}


func searchArtist(w http.ResponseWriter, r *http.Request) {

	jsonOut, _ := json.Marshal(Song{
				Song : "sweet child of mine",
				Artist : "guns and roses",
				Genre : "rock",
				Lenght : 503,
			})
    	fmt.Fprintf(w, string(jsonOut))

    
}
func logging(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %v\n", r.URL)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }
