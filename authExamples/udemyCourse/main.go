package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

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
	r.HandleFunc("/protected", tokenVerifyMiddleWare(protectedEndpoint)).Methods("GEt")

	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func respondWithError(w http.ResponseWriter, status int, errMsg string) {

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Error{errMsg})
}

func respondWithJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func decodeUserHttp(r *http.Request) User {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	return user
}

func handledBadUser(w http.ResponseWriter, user User) bool {
	if user.Email == "" {
		respondWithError(w, http.StatusBadRequest, "The email is missing")
		return true
	}
	if user.Password == "" {
		respondWithError(w, http.StatusBadRequest, "The password is missing")
		return true
	}
	return false
}

func generateToken(user User) (string, error) {
	// in the production system take this as an environment variable or command line arg
	const secret = "secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func signup(w http.ResponseWriter, r *http.Request) {
	user := decodeUserHttp(r)
	if handledBadUser(w, user) {
		return
	}

	hashWord, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to hash password")
		return
	}
	user.Password = string(hashWord)
	sql := "insert into users (email, password) values ($1, $2) RETURNING id;"
	err = db.QueryRow(sql, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to insert into db")
	}
	user.Password = ""
	respondWithJson(w, user)
}

func login(w http.ResponseWriter, r *http.Request) {
	user := decodeUserHttp(r)
	if handledBadUser(w, user) {
		return
	}

	incomingPw := user.Password

	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusBadRequest, "the user does not exist")
		} else {
			fmt.Println(err)
			respondWithError(w, http.StatusInternalServerError, "some problem with the db")
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(incomingPw))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid password")
	}

	token, err := generateToken(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "some problem with the token")
	}
	respondWithJson(w, JWT{token})
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProtectedEndpoint")
}

func tokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tmpSlice := strings.Split(authHeader, " ")

		if len(tmpSlice) != 2 {
			respondWithError(w, http.StatusUnauthorized, "authHeader not in expected format")
			return
		}

		authToken := tmpSlice[1]

		token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("There was an error")
			}

			// note here is the secret. don't put this in source control for real system
			return []byte("secret"), nil
		})

		if error != nil {
			respondWithError(w, http.StatusUnauthorized, error.Error())
			return
		}

		if !token.Valid {
			respondWithError(w, http.StatusUnauthorized, "token is not valid")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func logFatalErrorOrNothing(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
