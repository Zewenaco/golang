package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test-api/pkg/models"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	var foundBook *models.Book
	w.Header().Add("Content-Type", "application/json")
	parameters := mux.Vars(r)
	strId := strings.TrimSpace(parameters["id"])

	if len(strId) == 0 {
		log.Println("There is missing require parameter id")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There is missing require parameter id")
		return
	}

	id, _ := strconv.Atoi(strId)
	/*
		for _, book := range mocks.Books {

			if book.Id == id {
				foundBook = &book
				break
			}
		}
		if foundBook == nil {
			log.Println("There is no existent book with id " + strId)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w)
			return
		}
	*/
	if result := h.DB.First(&foundBook, id); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("There is no existent book with id " + strId)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundBook)
}
