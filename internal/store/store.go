package store

import (
	"github.com/ahmadmilzam/go/internal/entity"
	"github.com/ahmadmilzam/go/pkg/pgclient"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// type TransferTxParams struct {
// 	FromAccountID int64 `json:"src_account_id"`
// 	ToAccountID   int64 `json:"dst_account_id"`
// 	Amount        int64 `json:"amount"`
// }

// // TransferTxResult is the result of the transfer transaction
// type TransferTxResult struct {
// 	Transfer    entity.Transfer `json:"journal"`
// 	FromAccount entity.Account  `json:"src_account"`
// 	ToAccount   entity.Account  `json:"dst_account"`
// 	SrcTransfer entity.Entry    `json:"src_transfer"`
// 	DstTransfer entity.Entry    `json:"dst_transfer"`
// }

type Store interface {
	entity.StoreQuerier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*sqlx.DB
	*Queries
}

// NewStore creates a new store
func NewStore() Store {
	sql := pgclient.New()

	return &SQLStore{
		DB:      sql,
		Queries: NewQueries(sql),
	}
	// alt version
	// return &Store{
	// 	AccountStore: &AccountStore{DB: db},
	// }, nil
}

// alt version
// type Store struct {
// 	*AccountStore // TODO: add another store here and in model.Store interface
// }
