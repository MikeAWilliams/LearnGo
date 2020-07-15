package httpdriver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
	"github.com/gorilla/mux"
)

func errorWasHandled(w http.ResponseWriter, err error) bool {
	if nil != err {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return true
	}
	return false
}

func getSlashHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func getGetItemsHandler(db busineslogic.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := busineslogic.GetAllItems(db)
		if errorWasHandled(w, err) {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}
}

func getGetSpecificItemHandler(db busineslogic.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		item, err := busineslogic.GetItem(title, db)
		if errorWasHandled(w, err) {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}

type postItemBody struct {
	Description string
}

func getPostSpecificItemHandler(db busineslogic.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		var body postItemBody
		body.Description = "not set"
		bodyError := json.NewDecoder(r.Body).Decode(&body)
		if errorWasHandled(w, bodyError) {
			return
		}

		added, item, err := busineslogic.AddItem(title, body.Description, db)
		if errorWasHandled(w, err) {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if !added {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
		}
		json.NewEncoder(w).Encode(item)
	}
}

func Start(db busineslogic.Database) {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", getSlashHandler).Methods("GET")
	r.HandleFunc("/api/v1/items", getGetItemsHandler(db)).Methods("GET")
	r.HandleFunc("/api/v1/items/{title}", getGetSpecificItemHandler(db)).Methods("GET")
	r.HandleFunc("/api/v1/items/{title}", getPostSpecificItemHandler(db)).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
