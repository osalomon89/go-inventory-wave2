package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/osalomon89/go-inventory-wave2/domain"
)

var inventory []domain.Book
var id uint

func main() {
	id = 4

	book1 := domain.Book{
		ID: 1,
		Author: domain.Author{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
		Stock: 9,
	}

	book2 := domain.Book{
		ID: 2,
		Author: domain.Author{
			Firstname: "Stanislaw",
			Lastname:  "Lem",
		},
		Title: "Solaris",
		Price: 65.20,
		ISBN:  "0156027607",
		Stock: 15,
	}

	book3 := domain.Book{
		ID: 3,
		Author: domain.Author{
			Firstname: "Arthur C.",
			Lastname:  "Clarck",
		},
		Title: "Rendezvous with Rama",
		Price: 53.50,
		ISBN:  "0-575-01587-X",
	}

	book4 := domain.Book{
		ID: 4,
		Author: domain.Author{
			Firstname: "Jorge Luis",
			Lastname:  "Borges",
		},
		Title: "El Aleph",
		Price: 42.75,
		ISBN:  "84-206-1933-7",
	}

	inventory = []domain.Book{
		book1,
		book2,
		book3,
		book4,
	}

	router := mux.NewRouter()

	const port string = ":8888"

	router.HandleFunc("/ping", ping).Methods("GET")

	router.HandleFunc("/inventory", listBooks).Methods("GET")
	router.HandleFunc("/inventory", addBook).Methods("POST")
	router.HandleFunc("/inventory/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/inventory/{id}", getBookByID).Methods("GET")

	log.Println("Server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

type ResponseInfo struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   "pong",
	})
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	isbn := r.URL.Query().Get("isbn")
	if isbn != "" {
		var sliceBooks []domain.Book
		for _, v := range inventory {
			if v.ISBN == isbn {
				sliceBooks = append(sliceBooks, v)
			}
		}

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: 200,
			Data:   sliceBooks,
		})
		return
	}

	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   inventory,
	})
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newBook domain.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusInternalServerError,
			Data:   err,
		})
		return
	}

	id = id + 1
	newBook.ID = id

	inventory = append(inventory, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusCreated,
		Data:   newBook,
	})
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	idParam := param["id"]

	idBook, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error: " + idParam,
		})
		return
	}

	var newBook domain.Book
	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusInternalServerError,
			Data:   err,
		})
		return
	}

	for i, book := range inventory {
		if book.ID == uint(idBook) {
			newBook.ID = book.ID
			inventory[i] = newBook
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   newBook,
	})
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	idParam := param["id"]

	idBook, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error: " + idParam,
		})
		return
	}

	var newBook domain.Book
	for _, book := range inventory {
		if book.ID == uint(idBook) {
			newBook = book
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   newBook,
	})
}
