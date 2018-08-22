package store

import "time"

type Transaction struct {
	Hash        string `sql:",pk"`
	From        string
	To          string
	Seen        bool
	BlockNumber int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
