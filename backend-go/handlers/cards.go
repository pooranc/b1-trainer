package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pooranc/b1-trainer/backend-go/db"
	"github.com/pooranc/b1-trainer/backend-go/models"
)

func GetAllCards(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, type, question, answer, hint, created_at FROM  cards")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cards []models.Card
	for rows.Next() {
		var c models.Card
		err := rows.Scan(&c.ID, &c.Type, &c.Question, &c.Answer, &c.Hint, &c.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cards = append(cards, c)
	}

	w.Header().Set("Contect-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DB.QueryRow(
		"INSERT INTO cards (type, question, answer, hint) VALUES ($1, $2, $3, $4) RETURNING id, created_at",
		card.Type, card.Question, card.Answer, card.Hint,
	).Scan(&card.ID, &card.CreatedAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("DELETE FROM cards WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetCardsByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardType := vars["type"]

	rows, err := db.DB.Query(
		"SELECT id, type, question, answer, hint, created_at FROM cards WHERE type = $1",
		cardType,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cards []models.Card
	for rows.Next() {
		var c models.Card
		err := rows.Scan(&c.ID, &c.Type, &c.Question, &c.Answer, &c.Hint, &c.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cards = append(cards, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

func RegisterCardRoutes(r *mux.Router) {
	r.HandleFunc("/api/cards", GetAllCards).Methods("GET")
	r.HandleFunc("/api/cards", CreateCard).Methods("POST")
	r.HandleFunc("/api/cards/{id}", DeleteCard).Methods("DELETE")
	r.HandleFunc("/api/cards/type/{type}", GetCardsByType).Methods("GET")
}
