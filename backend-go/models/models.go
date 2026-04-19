package models

import "time"

type Card struct {
    ID        int64     `json:"id"`
    Type      string    `json:"type"`
    Question  string    `json:"question"`
    Answer    string    `json:"answer"`
    Hint      string    `json:"hint"`
    CreatedAt time.Time `json:"createdAt"`
}

type UserCardProgress struct {
    ID             int64     `json:"id"`
    Card           Card      `json:"card"`
    CardID         int64     `json:"cardId"`
    Repetition     int       `json:"repetition"`
    Easiness       float64   `json:"easiness"`
    IntervalDays   int       `json:"intervalDays"`
    NextReviewDate time.Time `json:"nextReviewDate"`
}