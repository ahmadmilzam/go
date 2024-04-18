package usecase

const (
	AccountRoleUnregistered = "UNREGISTERED"
	AccountRoleRegistered   = "REGISTERED"
	AccountRoleInternalCoa  = "INTERNAL_COA"

	AccountStatusPending  = "PENDING"
	AccountStatusActive   = "ACTIVE"
	AccountStatusInactive = "INACTIVE"
	AccountStatusBlocked  = "BLOCKED"

	WalletTypeCash  = "CASH"
	WalletTypePoint = "POINT"

	TransferTypeDefault    = "TRANSFER"
	TransferTypeTopup      = "TOPUP"
	TransferTypePayment    = "PAYMENT"
	TransferTypeReversal   = "REVERSAL"
	TransferTypeCorrection = "CORRECTION"

	CurrencyIDR = "IDR"
)

func GetSupportedAccountRole() []string {
	return []string{
		AccountRoleUnregistered,
		AccountRoleRegistered,
		AccountRoleInternalCoa,
	}
}

func GetSupportedTransferType() []string {
	return []string{
		TransferTypeDefault,
		TransferTypeTopup,
		TransferTypePayment,
		TransferTypeReversal,
		TransferTypeCorrection,
	}
}

func GetSupportedAccountStatus() []string {
	return []string{
		AccountStatusPending,
		AccountStatusActive,
		AccountStatusInactive,
		AccountStatusBlocked,
	}
}
