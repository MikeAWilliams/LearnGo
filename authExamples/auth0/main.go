// following the example from https://auth0.com/blog/authentication-in-golang/
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// apis
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/products", ProductsHandler).Methods("GET")
	r.Handle("/products/{slug}/feedback", AddFeedbackHandler).Methods("POST")

	// Our application will run on port 8080. Here we declare the port and pass in our router.
	http.ListenAndServe(":8080", r)
}
