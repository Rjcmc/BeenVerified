package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"goji.io"
	"goji.io/pat"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	
)

type Song struct {
	Song   string  "json:song"
	Artist  string  "json:artist"
	Genre   string  "json:genre"
	length  int     "json:length"
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/artist/:artistName"), searchArtist)
	mux.HandleFunc(pat.Get("/song/:songName"), searchSong)
	mux.Use(logging)
	http.ListenAndServe("localhost:8000", mux)
}

func Connection(w http.ResponseWriter, where string)(error){
    db, err := sql.Open("sqlite3", "./jrdd.db")
    checkErr(err)
    rows, err := db.Query("SELECT songs.artist,songs.song,genres.name,songs.length FROM genres inner join songs on genres.id=songs.genre where "+where )
    checkErr(err)
    var artist string
    var song   string
    var name   string
    var length int
    list:= []Song{}
	
    for rows.Next() {
        err = rows.Scan(&artist, &song,&name,&length)
        checkErr(err)
	b:=Song{
		Artist:    artist,
		Song:      song,
		Genre:     name,
		length:    length,
		}
        list=append(list,b)  // append to the list of songs        
    } 
    jsonOut, _ := json.Marshal(list)
    fmt.Fprintf(w, string(jsonOut))
    rows.Close() 
    db.Close()
    return err

}

func searchArtist(w http.ResponseWriter, r *http.Request) {

	artistName := pat.Param(r, "artistName")
	Connection(w ,"songs.artist='"+artistName+"'" )// call the conection to the data base with the where condition for an artist search

    
}

func searchSong(w http.ResponseWriter, r *http.Request) {

	song := pat.Param(r, "songName")
        Connection(w,"songs.song='"+song+"'")// call the conection to the data base with the where condition for a song search

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
