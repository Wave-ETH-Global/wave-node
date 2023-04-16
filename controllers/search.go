package controllers

import (
	"github.com/Wave-ETH-Global/wave-node/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type SearchController struct {
	db *sqlx.DB
}

func NewSearchController(db *sqlx.DB) *SearchController {
	return &SearchController{db: db}
}

func (sc *SearchController) Search(ctx echo.Context) error {
	var acquantancies []models.Profile
	var req struct {
		Tags []string `json:"tags"`
	}
	ctx.Bind(&req)
	sc.db.Get(acquantancies, `
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
	`, ctx.QueryParam("uuid"), pq.Array(req.Tags))

	return nil
}
