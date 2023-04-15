package models

type Profile struct {
	UUID       string                   `db:"uuid"`
	ETHAddress string                   `db:"eth_address"`
	Username   string                   `db:"username"`
	Metadata   map[string]interface{}   `db:"metadata"`
	PublicTags []string                 `db:"public_tags"`
	Tokens     []map[string]interface{} `db:"tokens"`
}

type ConnectionStatus string

const (
	Requested ConnectionStatus = "requested"
	Connected ConnectionStatus = "connected"
	Denied    ConnectionStatus = "denied"
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
