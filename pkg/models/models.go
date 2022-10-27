package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Pin struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Created time.Time 
	Expires time.Time 
}