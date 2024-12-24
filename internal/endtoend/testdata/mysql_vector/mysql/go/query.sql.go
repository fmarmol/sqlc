// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"
)

const insertVector = `-- name: InsertVector :exec
INSERT INTO foo(embedding) VALUES (STRING_TO_VECTOR('[0.1, 0.2, 0.3, 0.4]'))
`

func (q *Queries) InsertVector(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, insertVector)
	return err
}

const selectVector = `-- name: SelectVector :many
SELECT id FROM foo
ORDER BY DISTANCE(STRING_TO_VECTOR('[1.2, 3.4, 5.6]'), embedding, 'L2_squared')
LIMIT 10
`

func (q *Queries) SelectVector(ctx context.Context) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, selectVector)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
