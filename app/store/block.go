package store

import "time"

type Block struct {
	Number       int64          `sql:",pk"`
	Transactions []*Transaction `sql:",fk:block_number"`
	BlockTime    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
