package models

import "time"

type Blog struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CoverImage string    `json:"coverimage"`
	CreatedAt  time.Time `json:"createdat"`
	ReadCount  int       `json:"readcount"`
	Published  bool      `json:"published"`
	Featured   bool      `json:"featured"`
	WrittenBy  uint      `json:"writtenby"`
	MetaTitle  string    `json:"metatitle"`
}
