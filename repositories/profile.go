package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/database"
	"github.com/Wave-ETH-Global/wave-node/models"
	"github.com/google/logger"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	AirStackURL = "https://api.airstack.xyz/gql"
)

type ProfileRepository struct {
	db  *sqlx.DB
	cfg *config.Config
}

func NewProfileRepository(db *sqlx.DB, cfg *config.Config) *ProfileRepository {
	return &ProfileRepository{
		db:  db,
		cfg: cfg,
	}
}

func (pr *ProfileRepository) GetOnchainInfo(ethAddress string, cursor string) (map[string]interface{}, error) {
	client := &http.Client{}

	query := fmt.Sprintf(`
	query QB5 {
		TokenBalances(input: {filter: { owner: {_eq: "%s"}}, limit: 50, blockchain: ethereum, cursor: "%s"}) {
			TokenBalance {
				chainId
				tokenAddress
      			tokenId
				tokenType
				token {
					name
					symbol
				}
			}
			pageInfo {
				nextCursor
			}
		}
	}`, ethAddress, cursor)

	body := map[string]string{
		"operationName": "QB5",
		"query":         query,
	}

	b, _ := json.Marshal(body)
	br := strings.NewReader(string(b))
	req, err := http.NewRequest(
		"POST", AirStackURL, br,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", pr.cfg.AirStackToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(respBytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type RequestConnectionParam struct {
	ConnectorAddress string
	RequestedAddress string
	Tags             []string
	PrivateTags      []string
}

func (pr *ProfileRepository) RequestConnection(ctx context.Context, param RequestConnectionParam) error {
	db := database.ProvideDatabase(pr.cfg)
	if db == nil {
		return nil
	}

	var requesterUUID string
	var requestedUUID string
	sel := `select uuid from profile where eth_address=$1`
	if err := db.QueryRow(sel, param.ConnectorAddress).Scan(&requesterUUID); err != nil {
		return errors.New("failed to get")
	}

	if err := db.QueryRow(sel, param.ConnectorAddress).Scan(&requestedUUID); err != nil {
		return errors.New("failed to get")
	}

	query := `INSERT INTO connection (vertex_a, vertex_b, status, tags, private_tags) VALUES ($1, $2, $3, $4, $5)`
	if err := db.MustExec(query, &requesterUUID, &requestedUUID, models.Requested, pq.Array(param.Tags), pq.Array(param.PrivateTags)); err != nil {
		logger.Error(err)
		return errors.New("failed to insert")
	}

	return nil
}
