package domain

import "context"

type Account struct {
	userName        string
	walletAddresses []Address
}

func NewAccount(userName string, walletAddresses []string) Account {
	addresses := make([]Address, 0, len(walletAddresses))

	for _, wa := range walletAddresses {
		addresses = append(addresses, Address(wa))
	}

	return Account{
		userName:        userName,
		walletAddresses: addresses,
	}
}

func (a *Account) WalletAddresses() []Address { return a.walletAddresses }
func (a *Account) Name() string               { return a.userName }

type AccountRepository interface {
	GetAccountByAddress(ctx context.Context, walletAddress Address) (*Account, error)
}
