package controllers

import "github.com/Wave-ETH-Global/wave-node/controllers/domain"

type HTTPError struct {
	Error string `json:"error"`
}

func walletAddressesToString(walletAddresses []domain.Address) []string {
	addresses := make([]string, 0, len(walletAddresses))

	for _, wa := range walletAddresses {
		addresses = append(addresses, string(wa))
	}

	return addresses
}
