package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/go/internal/entity"
	"github.com/lib/pq"
)

const (
	createAccountSQL = `
	INSERT INTO accounts
	VALUES(:phone, :name, :email, :role, :status, :created_at, :updated_at)`
	updateAccountSQL = `
	UPDATE accounts
	SET
		phone = :phone,
		name = :name,
		email = :email,
		role = :role,
		status = :status,
		updated_at = :updated_at
	WHERE phone = :phone`
	deleteAccountByIdSQL         = `DELETE * FROM accounts WHERE phone = :phone`
	findAccountForUpdateByIdSQL  = `SELECT * FROM accounts WHERE phone = $1 LIMIT 1 FOR UPDATE`
	findAccountAndWalletsByIdSQL = `
	SELECT
		account.*,
		wallet.id "wallet.id",
		wallet.type "wallet.type",
		wallet.balance "wallet.balance",
		wallet.created_at "wallet.created_at",
		wallet.updated_at "wallet.updated_at"
	FROM
		accounts AS account
		JOIN wallets AS wallet ON account.phone = wallet.account_phone
	WHERE
		account.phone = $1`
)

func (s *Queries) CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	_, err := s.db.NamedExecContext(ctx, createAccountSQL, account)
	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
		fmt.Println("pq error:", err)
		fmt.Println("pq error:", err.Code.Name())
		/*
			pq error: pq: duplicate key value violates unique constraint "accounts_pkey"
			pq error: unique_violation
		*/
	}
	if err != nil {
		err = fmt.Errorf("CreateAccount: %w", err)
		return nil, err
	}

	return account, nil
}

func (s *SQLStore) CreateAccountTx(ctx context.Context, account *entity.Account, wallets []entity.Wallet, counter *entity.TransferCounter) error {

	err := s.execTx(func(q *Queries) error {
		var err error

		_, err = q.CreateAccount(ctx, account)
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		_, err = q.CreateWallet(ctx, &wallets[0])
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		_, err = q.CreateWallet(ctx, &wallets[1])
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		_, err = q.CreateCounter(ctx, counter)
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		return err
	})

	return err
}

func (s *Queries) UpdateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	_, err := s.db.NamedExecContext(ctx, updateAccountSQL, account)

	if err != nil {
		return nil, fmt.Errorf("UpdateAccount: %w", err)
	}

	return account, nil
}

func (s *Queries) DeleteAccount(ctx context.Context, id string) error {
	_, err := s.db.NamedExecContext(ctx, deleteAccountByIdSQL, id)
	if err != nil {
		return fmt.Errorf("DeleteAccount: %w", err)
	}

	return nil
}

func (s *Queries) FindAccountForUpdateById(ctx context.Context, phone string) (*entity.Account, error) {
	a := &entity.Account{}
	err := s.db.GetContext(ctx, a, findAccountForUpdateByIdSQL, phone)
	if err != nil {
		return nil, fmt.Errorf("FindAccountForUpdateByPhone: %w", err)
	}

	return a, nil
}

func (s *Queries) FindAccountById(ctx context.Context, phone string) (*entity.Account, error) {
	ma := &entity.Account{}
	err := s.db.GetContext(ctx, ma, `SELECT * FROM accounts WHERE phone = $1 LIMIT 1`, phone)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
			fmt.Println("pq error:", err)
			fmt.Println("pq error:", err.Code.Name())
			/*
				pq error: pq: duplicate key value violates unique constraint "accounts_pkey"
				pq error: unique_violation
			*/
		}
		err = fmt.Errorf("FindAccountByPhone: %w", err)
		return nil, err
	}

	return ma, nil
}

func (s *Queries) FindAccountAndWalletsById(ctx context.Context, phone string) ([]entity.AccountWallet, error) {
	var accWallets []entity.AccountWallet

	err := s.db.SelectContext(ctx, &accWallets, findAccountAndWalletsByIdSQL, phone)
	if err != nil {
		err = fmt.Errorf("FindAccountAndWalletsById: %w", err)
		return nil, err
	}

	return accWallets, nil
}
