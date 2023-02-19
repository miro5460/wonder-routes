package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostForm = %v\n", r.PostForm)
		name := r.FormValue("name")
		fmt.Fprintf(w, "Name= %s\n", name)
	default:
		fmt.Fprint(w, "Method not supported.")
	}
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

// func access_strava(w http.ResponseWriter, r. *http.Request) {

// }

func main() {
	http.HandleFunc("/", hello)
	//http.HandleFunc("/access_strava", access_strava)

	fmt.Printf("Starting HTTP server...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
