package repositories

import (
	"context"
	"log"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/controllers/domain"
	"github.com/Wave-ETH-Global/wave-node/database"
	"github.com/Wave-ETH-Global/wave-node/models"
	"github.com/google/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AccountRepo struct {
	db  *sqlx.DB
	cfg *config.Config
}

type QueryLogger struct {
	queryer sqlx.Queryer
	logger  *log.Logger
}

func NewAccountRepository(
	db *sqlx.DB,
	cfg *config.Config,
) *AccountRepo {
	return &AccountRepo{
		db:  db,
		cfg: cfg,
	}
}

func (r *AccountRepo) GetAccountByAddress(
	ctx context.Context,
	walletAddress domain.Address,
) *domain.Account {
	db := database.ProvideDatabase(r.cfg)
	if db == nil {
		return nil
	}

	p := models.Profile{
		ETHAddress: string(walletAddress),
	}
	query := `SELECT uuid, eth_address FROM profile WHERE eth_address=:eth_address`
	_, err := db.NamedExec(query, p)
	if err != nil {
		logger.Error(err)
		return nil
	}

	profile := domain.NewAccount("", []string{string(walletAddress)})
	return &profile
}

func (r *AccountRepo) CreateAddress(
	ctx context.Context,
	walletAddress domain.Address,
) *domain.Account {
	db := database.ProvideDatabase(r.cfg)

	if db == nil {
		return nil
	}

	query := `INSERT INTO profile (uuid, eth_address) VALUES (:uuid, :eth_address)`
	prof := models.Profile{
		UUID:       uuid.New().String(),
		ETHAddress: string(walletAddress),
	}
	_, err := db.NamedExec(query, prof)
	if err != nil {
		return nil
	}

	profile := domain.NewAccount("", []string{string(walletAddress)})
	return &profile
}
