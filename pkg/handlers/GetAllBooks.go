package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-api/pkg/models"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("There is no record found")
		return
	}
	json.NewEncoder(w).Encode(books)
}
