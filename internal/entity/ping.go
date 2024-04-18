package entity

import (
	"context"
	"time"
)

//go:generate mockery --name AccountStoreQuerier
type AccountQuery interface {
	CreateAccount(ctx context.Context, account *Account) (*Account, error)
	CreateAccountTx(ctx context.Context, account *Account, wallets []Wallet, counter *TransferCounter) error
	UpdateAccount(ctx context.Context, account *Account) (*Account, error)
	FindAccountForUpdateById(ctx context.Context, id string) (*Account, error)
	FindAccountById(ctx context.Context, id string) (*Account, error)
	FindAccountAndWalletsById(ctx context.Context, id string) ([]AccountWallet, error)
}

type Account struct {
	Phone     string    `db:"phone" faker:"customphone,unique"`
	Name      string    `db:"name" faker:"name,unique"`
	Email     string    `db:"email" faker:"email,unique"`
	Role      string    `db:"role" faker:"accountRole"`
	Status    string    `db:"status" faker:"accountStatus"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AccountWallet struct {
	Account
	Wallet `db:"wallet"`
}
