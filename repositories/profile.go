package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/jmoiron/sqlx"
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
