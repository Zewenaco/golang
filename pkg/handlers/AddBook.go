package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"test-api/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("There is a problem with the body " + string(body))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
	}

	var book models.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Println("There is a problem with the body " + string(body))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
	}

	// Append to the Book mocks
	/*book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)*/
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("There is conflict to create new record")
		return
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
