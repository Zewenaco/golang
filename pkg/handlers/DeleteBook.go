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

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	parameters := mux.Vars(r)
	strId := strings.TrimSpace(parameters["id"])

	if len(strId) == 0 {
		log.Println("There is missing require parameter id")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There is missing require parameter id ")
		return
	}

	id, _ := strconv.Atoi(strId)
	/*
		var index *int
		for i, book := range mocks.Books {
			if book.Id == id {
				index = &i
				break
			}
		}
		 if index == nil {
				log.Println("There is no record found with id " + strId)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("There is no record found with id " + strId)
				return
		}
		/*
			ret := make([]models.Book, 0)
			ret = append(ret, mocks.Books[:*index]...)
			mocks.Books = append(ret, mocks.Books[*index+1:]...)
	*/
	var book *models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("There is no record found with id " + strId)
		return
	}
	h.DB.Delete(&book)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
