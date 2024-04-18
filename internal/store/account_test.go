package store

import (
	"context"
	"testing"
	"time"

	"github.com/ahmadmilzam/go/internal/entity"
	"github.com/ahmadmilzam/go/pkg/randomizer"
	"github.com/ahmadmilzam/go/pkg/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountTx(t *testing.T) {
	a := &entity.Account{}
	err := randomizer.RandomAccountData(a)
	require.NoError(t, err)

	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now

	wc := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: a.Phone,
		Balance:      0.00,
		Type:         "CASH",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	wp := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: a.Phone,
		Balance:      0.00,
		Type:         "POINT",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	ww := []entity.Wallet{}

	ww = append(ww, wc, wp)

	tc := &entity.TransferCounter{
		WalletId:            wc.ID,
		CreditCountDaily:    0,
		CreditCountMonthly:  0,
		CreditAmountDaily:   0,
		CreditAmountMonthly: 0,
		CreatedAt:           now,
		UpdatedAt:           now,
	}

	err1 := testStore.CreateAccountTx(context.Background(), a, ww, tc)

	require.NoError(t, err1)
	// require.Equal(t, d.Phone, ac.Phone)
	// require.Equal(t, d.Name, ac.Name)
	// require.Equal(t, d.Email, ac.Email)
	// require.Equal(t, d.Role, ac.Role)
	// require.Equal(t, d.Status, ac.Status)
}
