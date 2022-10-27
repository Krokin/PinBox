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
	Created time.Time `json:"created"`
	Expires time.Time `json:"expires"`
}