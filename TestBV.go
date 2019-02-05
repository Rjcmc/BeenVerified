package main

import (
	
	"fmt"
	"net/http"
	"goji.io"
	"goji.io/pat"
	
)


func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/artist/:artistName"), searchArtist)
	mux.Use(logging)
	http.ListenAndServe("localhost:8000", mux)
}


func searchArtist(w http.ResponseWriter, r *http.Request) {

	artistName := pat.Param(r, "artistName")
	w.Write([]byte(fmt.Sprintf("Hello %s!",artistName)))

    
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
