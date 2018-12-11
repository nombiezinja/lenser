package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")

// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Placeholder</h1>")
// 	} else if r.URL.Path == "/contact" {
// 		fmt.Fprint(w, "Reach out at "+"<a href= \"mailto: tianyizhang1987@gmail.com\">"+"tianyizhang1987@gmail.com</a>")
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>We could not find the page "+"were looking for:(</h1>"+"<p> Please email us if you keep being sent to an invalid page </p>")
// 	}
// }

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Placeholder</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Reach out at "+"<a href= \"mailto: tianyizhang1987@gmail.com\">"+"tianyizhang1987@gmail.com</a>")
}
