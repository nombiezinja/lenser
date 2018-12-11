package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	router.NotFound = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", router)
}

func home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Placeholder</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Reach out at "+"<a href= \"mailto: tianyizhang1987@gmail.com\">"+"tianyizhang1987@gmail.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ</h1>"+"<p>Placeholder placeholder placeholder</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Page not found.</h1>")
}
