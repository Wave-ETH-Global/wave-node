package domain

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/logger"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Address string

const AddressMaxLength = 42

var (
	ErrAddressInvalid = errors.New("address: must start with '0x'")
	ErrAddressTooLong = fmt.Errorf(
		"address: must be up to %d characters",
		AddressMaxLength,
	)
	ErrSigVerifyECRecoverFailed    = errors.New("address: ecrecover failed")
	ErrSigVerifyAddressMismatch    = errors.New("address: address mismatch")
	ErrSigVerifyVerificationFailed = errors.New("address: signature verification failed")
	ErrSigVerifyFailed             = errors.New("address: verify signature failed")
)

func NewAddress(v string) (Address, error) {
	if len([]rune(v)) > AddressMaxLength {
		return "", errors.WithStack(ErrAddressTooLong)
	}

	if !strings.HasPrefix(v, "0x") {
		return "", errors.WithStack(ErrAddressInvalid)
	}

	return Address(v), nil
}

func (a Address) SigVerify(message string, signature string) error {
	sig, err := hexutil.Decode(string(signature))
	if err != nil {
		logger.Errorf("<SigVerify> public key mismatch. address:%s, message:%s, signature:%s", a, message, signature)
		return errors.WithStack(ErrSigVerifyFailed)
	}

	pubKey, msgHash, err := a.ecrecover(message, sig)
	if err != nil {
		logger.Errorf("<SigVerify> invalid signature format. address:%s, message:%s, sig:%+v", a, message, sig)
		return errors.WithStack(ErrSigVerifyECRecoverFailed)
	}

	pubAddress := crypto.PubkeyToAddress(*pubKey)
	log.Info().Msgf("<SigVerify> recovered address. pubAddress:%s", pubAddress)

	pubKeyAddress, err := NewAddress(pubAddress.Hex())
	if err != nil {
		logger.Errorf("<SigVerify> public key convert to address error!! err:%+v", err)
		return errors.WithStack(ErrSigVerifyFailed)
	}
	if !pubKeyAddress.Equal(a) {
		logger.Errorf("<SigVerify> public key mismatch. expected:%s, recovered:%s, message:%s, signature:%s", a, pubKeyAddress, message, signature)
		return errors.WithStack(ErrSigVerifyAddressMismatch)
	}

	sigWithoutRecoverID, err := a.trimRecoverID(sig)
	if err != nil {
		logger.Errorf("<SigVerify> trim recover id error. address:%s, len(sig):%d, sig:%+v, err:%+v", pubKeyAddress, len(sig), sig, err)
		return errors.WithStack(ErrSigVerifyFailed)
	}
	log.Info().Msgf("<SigVerify> to verify... address:%s, sigWithoutRecoverID:%s, len(msgHash):%d, len(sig):%d", pubKeyAddress, sigWithoutRecoverID, len(msgHash), len(sig))

	if !crypto.VerifySignature(crypto.FromECDSAPub(pubKey), msgHash, sig[:len(sig)-1]) {
		logger.Errorf("<SigVerify> can't verify signature with key. address:%s, message:%s, signature:%s", a, message, signature)
		return errors.WithStack(ErrSigVerifyVerificationFailed)
	}

	log.Info().Msgf("<SigVerify> verified signature with key. address:%s, recovered address: %s, message:%s, signature:%s", a, pubKeyAddress, message, signature)
	return nil
}

func (a Address) trimRecoverID(sig []byte) ([]byte, error) {
	if len(sig) != 65 {
		return nil, errors.WithStack(ErrSigVerifyFailed)
	}
	return sig[:len(sig)-1], nil
}

func (a Address) Equal(b Address) bool {
	return strings.ToLower(string(a)) == strings.ToLower(string(b))
}

func (a Address) Has(addresses []Address) bool {
	for _, ad := range addresses {
		if strings.ToLower(string(a)) == strings.ToLower(string(ad)) {
			return true
		}
	}
	return false
}

func (a Address) ecrecover(message string, sig []byte) (*ecdsa.PublicKey, []byte, error) {
	if len(sig) != 65 {
		logger.Errorf("<ecrecover> invalid signature format(length). address:%s, message:%s, sig:%+v", a, message, sig)
		return nil, nil, errors.WithStack(ErrSigVerifyECRecoverFailed)
	}

	if sig[64] != 27 && sig[64] != 28 {
		logger.Warningf("<ecrecover> invalid signature format(last byte). address:%s, message:%s, sig:%+v", a, message, sig)
	}

	// note: personal_sign
	// https://github.com/ethereum/go-ethereum/blob/master/internal/ethapi/api.go#L545_L559kjklj
	_, ethSignMessage := accounts.TextAndHash([]byte(message))
	msgHash := crypto.Keccak256([]byte(ethSignMessage))
	logger.Infof("<ecrecover> get the public key.... address:%s, message:%s, []byte(signature):%+v", a, message, sig)

	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	pubKey, err := crypto.SigToPub(msgHash, sig)
	if err != nil {
		logger.Errorf("<ecrecover> invalid signature format(public key could not be obtained). message:%s, msgHash:%s, sig:%+v", a, message, sig)
		return nil, nil, errors.WithStack(ErrSigVerifyECRecoverFailed)
	}

	return pubKey, msgHash, nil
}

const (
	WalletSignatureMessage = "Sign\n        \n        Nonce: %s"
)
