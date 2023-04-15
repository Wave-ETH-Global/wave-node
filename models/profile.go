package models

type Profile struct {
	UUID       string                   `db:"uuid"`
	ETHAddress string                   `db:"eth_address"`
	Username   string                   `db:"username"`
	Metadata   map[string]interface{}   `db:"metadata"`
	PublicTags []string                 `db:"public_tags"`
	Tokens     []map[string]interface{} `db:"tokens"`
}

type Connection struct {
	ID       int64    `db:"id"`
	VertexA  string   `db:"vertex_a"`
	VertexB  string   `db:"vertex_a"`
	TagsA    []string `db:"tags_a"`
	TagsB    []string `db:"tags_b"`
	ProfileA *Profile `db:"-"`
	ProfileB *Profile `db:"-"`
}
