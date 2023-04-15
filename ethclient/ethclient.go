package ethclient

import (
	"context"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/repositories/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	client   *ethclient.Client
	contract *contracts.WaveUserHandleContract
}

func NewEthClient(cfg *config.Config) (*EthClient, error) {
	c := &EthClient{}

	client, err := ethclient.Dial(cfg.Ethereum.GatewayAddress)
	if err != nil {
		return nil, err
	}
	c.client = client

	contract, err := contracts.NewWaveUserHandleContract(common.HexToAddress(cfg.Ethereum.ContractAddress), client)
	if err != nil {
		return nil, err
	}

	c.contract = contract
	return c, nil
}

func (ec *EthClient) GetUUIDOfClaimedUserHandle(userhandle string) (string, error) {
	callOpts := &bind.CallOpts{Context: context.TODO(), Pending: false}
	uuid, err := ec.contract.GetUUIDByUserHandle(callOpts, userhandle)
	if err != nil {
		return "", nil
	}

	return uuid, nil
}
