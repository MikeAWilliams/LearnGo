package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	pgUrl, err := pq.ParseURL(os.Args[1])
	logFatalErrorOrNothing(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatalErrorOrNothing(err)

	err = db.Ping()
	logFatalErrorOrNothing(err)

	r := mux.NewRouter()
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GEt")

	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func respondWithError(w http.ResponseWriter, status int, errMsg string) {

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Error{errMsg})
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	spew.Dump(user)

	if user.Email == "" {
		respondWithError(w, http.StatusBadRequest, "The email is missing")
		return
	}

	if user.Password == "" {
		respondWithError(w, http.StatusBadRequest, "The password is missing")
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProtectedEndpoint")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleWare")
	return nil
}

func logFatalErrorOrNothing(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
