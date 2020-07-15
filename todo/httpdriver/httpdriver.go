package httpdriver

import (
	"log"
	"net/http"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
	"github.com/gorilla/mux"
)

func getSlashHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func getGetItemsHandler(db busineslogic.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := busineslogic.GetAllItems(db)
		if nil != err {
			w.Write([]byte(err.Error()))
		}

		result := ""
		for _, item := range items {
			result += item.String() + "\n"
		}
		w.Write([]byte(result))
	}
}

func getGetSpecificItemHandler(db busineslogic.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		item, err := busineslogic.GetItem(title, db)
		if nil != err {
			w.Write([]byte(err.Error()))
		}

		result := item.String() + "\n"
		w.Write([]byte(result))
	}
}

func Start(db busineslogic.Database) {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", getSlashHandler).Methods("GET")
	r.HandleFunc("/api/v1/items", getGetItemsHandler(db)).Methods("GET")
	r.HandleFunc("/api/v1/items/{title}", getGetSpecificItemHandler(db)).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
