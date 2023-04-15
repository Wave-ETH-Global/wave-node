package models

import (
	"encoding/json"
	"errors"

	"github.com/lib/pq"
)

type Profile struct {
	UUID       string             `db:"uuid"`
	ETHAddress string             `db:"eth_address"`
	Username   string             `db:"username"`
	Metadata   StringInterfaceMap `db:"metadata"`
	PublicTags pq.StringArray     `db:"public_tags"`
	Tokens     ArrayInterfaceMap  `db:"tokens"`
}

type ConnectionStatus string

const (
	Requested ConnectionStatus = "requested"
	Connected ConnectionStatus = "connected"
	denied    ConnectionStatus = "denied"
)

type Connection struct {
	ID          int64            `db:"id"`
	VertexA     string           `db:"vertex_a"`
	VertexB     string           `db:"vertex_b"`
	Status      ConnectionStatus `db:"status"`
	Read        bool             `db:"read"`
	Tags        []string         `db:"tags"`
	PrivateTags []string         `db:"private_tags"`
	ProfileA    *Profile         `db:"-"`
	ProfileB    *Profile         `db:"-"`
}

type StringInterfaceMap map[string]interface{}

func (m *StringInterfaceMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("failed type assertion to []byte")
	}
	return json.Unmarshal(b, &m)
}

type ArrayInterfaceMap []map[string]interface{}

func (m *ArrayInterfaceMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("failed type assertion to []byte")
	}
	return json.Unmarshal(b, &m)
}
