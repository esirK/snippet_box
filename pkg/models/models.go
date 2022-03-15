package models

import "time"

type Snippet struct {
	ID      int				`json:"id"`
	Title   string			`json:"title"`
	Content string			`json:"content"`
	Created time.Time		`json:"created"`
	Expires time.Time		`json:"expires"`
}
