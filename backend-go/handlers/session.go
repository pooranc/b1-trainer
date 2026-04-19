package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pooranc/b1-trainer/backend-go/algorithm"
	"github.com/pooranc/b1-trainer/backend-go/db"
	"github.com/pooranc/b1-trainer/backend-go/models"
)

func GetDueCards(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
        SELECT p.id, p.repetition, p.easiness, p.interval_days, p.next_review_date,
               c.id, c.type, c.question, c.answer, c.hint
        FROM user_card_progress p
        JOIN cards c ON p.card_id = c.id
        WHERE p.next_review_date <= $1
    `, time.Now())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var progressList []models.UserCardProgress

	for rows.Next() {
		var p models.UserCardProgress
		var c models.Card

		err := rows.Scan(
			&p.ID, &p.Repetition, &p.Easiness, &p.IntervalDays, &p.NextReviewDate,
			&c.ID, &c.Type, &c.Question, &c.Answer, &c.Hint,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p.Card = c
		progressList = append(progressList, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(progressList)
}

type RatingRequest struct {
	CardID int64 `json:"cardId"`
	Rating int   `json:"rating"`
}

func SubmitRating(w http.ResponseWriter, r *http.Request) {
	var req RatingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var p models.UserCardProgress
	var c models.Card
	err = db.DB.QueryRow(`
        SELECT p.id, p.repetition, p.easiness, p.interval_days, p.next_review_date,
               c.id, c.type, c.question, c.answer, c.hint
        FROM user_card_progress p
        JOIN cards c ON p.card_id = c.id
        WHERE p.card_id = $1
    `, req.CardID).Scan(
		&p.ID, &p.Repetition, &p.Easiness, &p.IntervalDays, &p.NextReviewDate,
		&c.ID, &c.Type, &c.Question, &c.Answer, &c.Hint,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.Card = c

	result := algorithm.Calculate(p.Repetition, p.Easiness, p.IntervalDays, req.Rating)

	nextReview := time.Now().AddDate(0, 0, result.IntervalDays)

	_, err = db.DB.Exec(`
        UPDATE user_card_progress
        SET repetition = $1, easiness = $2, interval_days = $3, next_review_date = $4
        WHERE id = $5
    `, result.Repetition, result.Easiness, result.IntervalDays, nextReview, p.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Repetition = result.Repetition
	p.Easiness = result.Easiness
	p.IntervalDays = result.IntervalDays
	p.NextReviewDate = nextReview

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func RegisterSessionRoutes(r *mux.Router) {
	r.HandleFunc("/api/session/due", GetDueCards).Methods("GET")
	r.HandleFunc("/api/session/rate", SubmitRating).Methods("POST")
}
