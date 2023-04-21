package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/models"
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

// Airstack integration lives here
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

func (pr *ProfileRepository) GetProfileByAddress(ethAddress string) (*models.Profile, error) {
	p := models.Profile{}
	err := pr.db.Get(&p, "select * from profile where eth_address = $1", ethAddress)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (pr *ProfileRepository) GetProfileByUUID(uuid string) (*models.Profile, error) {
	p := models.Profile{}
	err := pr.db.Get(&p, "select * from profile where uuid = $1", uuid)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (pr *ProfileRepository) InsertProfile(p *models.Profile) error {
	mj, _ := json.Marshal(p.Metadata)
	tj, _ := json.Marshal(p.Tokens)
	_, err := pr.db.Exec("insert into profile (uuid, eth_address, username, public_tags, metadata, tokens) values ($1, $2, $3, $4, $5, $6)", p.UUID, p.ETHAddress, p.Username, pq.Array(p.PublicTags), mj, tj)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProfileRepository) GetProfileConnections(uuid string) ([]*models.Connection, error) {
	var list []*models.Connection
	err := pr.db.Select(list, "select * from connection where vertex_a = $1", uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*models.Connection{}, nil
		}
		return nil, err
	}

	profileA, err := pr.GetProfileByUUID(uuid)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		v.ProfileA = profileA
		v.ProfileB, err = pr.GetProfileByUUID(v.VertexB)
	}

	return list, nil
}

func (pr *ProfileRepository) SearchProfiles(uuid string, tags []string) ([]*models.Profile, error) {
	var acquantancies []*models.Profile
	err := pr.db.Select(&acquantancies, `
		WITH RECURSIVE traverse(vertex_a, vertex_b, tags, path, depth) AS (
    		SELECT
        		vertex_a,
        		vertex_b,
        		tags,
        		ARRAY[ROW(vertex_a, vertex_b)] AS path,
        		0 as depth
    		FROM
        		connection
    		WHERE
        		connection.vertex_a = $1
    		UNION ALL
    		SELECT connection.vertex_a, connection.vertex_b, connection.tags, path || ROW(connection.vertex_a, connection.vertex_b), traverse.depth+1
    		FROM traverse
    		JOIN
    		connection ON connection.vertex_a = traverse.vertex_b
    		WHERE ROW(connection.vertex_a, connection.vertex_b) <> ALL(path)
		)

		SELECT distinct on (vertex_b) profile.* from traverse
		join profile on uuid = vertex_b WHERE vertex_b <> $1 and traverse.depth < 3 and ($2 <@ traverse.tags or $2 <@ profile.public_tags);
	`, uuid, pq.Array(tags))

	if err != nil {
		if err == sql.ErrNoRows {
			return []*models.Profile{}, nil
		}
		return nil, err
	}

	return acquantancies, nil
}
