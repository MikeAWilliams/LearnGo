package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GEt")

	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
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
