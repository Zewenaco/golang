package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test-api/pkg/models"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var foundBook *models.Book
	w.Header().Add("Content-Type", "application/json")
	parameters := mux.Vars(r)
	strId := strings.TrimSpace(parameters["id"])

	if len(strId) == 0 {
		log.Println("There is missing require parameter id")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There is missing require parameter id ")
		return
	}

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("There is a problem with the body " + string(body))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
		return
	}

	var updatedBook models.Book
	err = json.Unmarshal(body, &updatedBook)
	if err != nil {
		log.Println("There is a problem with the body " + string(body))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
		return
	}

	id, _ := strconv.Atoi(strId)
	/*
		var index int
		for i, book := range mocks.Books {
			if book.Id == id {
				foundBook = &book
				index = i
				break
			}
		}
	*/
	if result := h.DB.First(&foundBook, id); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("There is no existent book with id " + strId)
		return
	}

	foundBook.Title = updatedBook.Title
	foundBook.Author = updatedBook.Author
	foundBook.Desc = updatedBook.Desc

	/*
		mocks.Books[index] = *foundBook
	*/
	h.DB.Save(&foundBook)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(foundBook)
}
