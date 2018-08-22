package store

import (
	"sync"

	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	Store struct {
		db          *pg.DB
		mu          *sync.Mutex
		confirms    int64
		blockNumber int64
	}
)

func checkDatabase(db *pg.DB) error {
	var err error
	// Check blocks table
	if _, err = db.ExecOne("SELECT 1 FROM blocks LIMIT 1"); err != nil && err != pg.ErrNoRows {
		return errors.WithMessage(err, "blocks table")
	}

	// Check transactions table
	if _, err = db.ExecOne("SELECT 1 FROM transactions LIMIT 1"); err != nil && err != pg.ErrNoRows {
		return errors.WithMessage(err, "transactions table")
	}

	// Check another

	return nil
}

func NewStore(db *pg.DB, v *viper.Viper) (*Store, error) {
	if err := checkDatabase(db); err != nil {
		return nil, errors.WithMessage(err, "store")
	}

	v.SetDefault("store.confirms", 6)

	s := &Store{
		db:       db,
		mu:       new(sync.Mutex),
		confirms: v.GetInt64("store.confirms"),
	}

	if err := s.lastBlockNumber(); err != nil {
		return nil, errors.WithMessage(err, "store")
	}

	return s, nil
}

// SetBlock current block number
func (s *Store) SetBlock(v int64) {
	s.mu.Lock()
	s.blockNumber = v
	s.mu.Unlock()
}

// Confirms is last confirmed block
func (s *Store) Confirms() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.blockNumber - s.confirms
}

// LastBlock in store
func (s *Store) LastBlock() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.blockNumber
}

// FetchTransactions to response
func (s *Store) FetchTransactions() (result []*Transaction, err error) {

	err = s.db.
		Model(&result).
		Where("block_number > ? OR seen = false", s.Confirms()).
		Order("block_number ASC").
		Select()

	return result, nil
}

// UpdateSeen for transactions
func (s *Store) UpdateSeen(txs []*Transaction) error {
	if len(txs) == 0 {
		return nil
	}

	_, err := s.db.
		Model(&txs).
		Column("seen").
		Update()

	return err
}

// SaveTransaction to store
func (s *Store) SaveTransaction(tx *Transaction) error {
	_, err := s.db.
		Model(tx).
		Where("hash = ?", tx.Hash).
		OnConflict("(hash) DO UPDATE").
		Set(`"from" = ?`, tx.From).
		Set(`"to" = ?`, tx.To).
		Set(`"block_number" = ?`, tx.BlockNumber).
		Set(`"updated_at" = NOW()`).
		SelectOrInsert()

	return err
}

// fetch last block number from database:
func (s *Store) lastBlockNumber() error {
	var num int64

	if _, err := s.db.QueryOne(&num, "SELECT MAX(number) FROM blocks LIMIT 1"); err != nil {
		return err
	}

	s.SetBlock(num)

	return nil
}

// SaveBlock to store
func (s *Store) SaveBlock(bl *Block) error {
	_, err := s.db.
		Model(bl).
		Where("number = ?", bl.Number).
		OnConflict("(number) DO UPDATE").
		Set("block_time = ?", bl.BlockTime).
		Set("updated_at = NOW()").
		SelectOrInsert()

	return err
}
