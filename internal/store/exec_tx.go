package store

import (
	"fmt"
)

func (s *SQLStore) execTx(fn func(*Queries) error) error {
	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	q := NewQueries(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
