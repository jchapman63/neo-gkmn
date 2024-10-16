// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: monsters.sql

package database

import (
	"context"
)

const fetchMonster = `-- name: FetchMonster :one
SELECT
    id, name, type, basehp
FROM
    monster
WHERE
    id = $1
`

func (q *Queries) FetchMonster(ctx context.Context, id string) (Monster, error) {
	row := q.db.QueryRow(ctx, fetchMonster, id)
	var i Monster
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.Basehp,
	)
	return i, err
}

const fetchMove = `-- name: FetchMove :one
SELECT
    id, name, power, type
FROM
    MOVE
WHERE
    id = $1
`

func (q *Queries) FetchMove(ctx context.Context, id string) (Move, error) {
	row := q.db.QueryRow(ctx, fetchMove, id)
	var i Move
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Power,
		&i.Type,
	)
	return i, err
}

const fetchStat = `-- name: FetchStat :one
SELECT
    monsterid, stattype, power
FROM
    stats
WHERE
    monsterID = $1::uuid
    AND statType = $2::text
`

type FetchStatParams struct {
	Column1 string
	Column2 string
}

func (q *Queries) FetchStat(ctx context.Context, arg FetchStatParams) (Stat, error) {
	row := q.db.QueryRow(ctx, fetchStat, arg.Column1, arg.Column2)
	var i Stat
	err := row.Scan(&i.Monsterid, &i.Stattype, &i.Power)
	return i, err
}

const listMonsters = `-- name: ListMonsters :many
SELECT
    id, name, type, basehp
FROM
    monster
`

func (q *Queries) ListMonsters(ctx context.Context) ([]Monster, error) {
	rows, err := q.db.Query(ctx, listMonsters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Monster
	for rows.Next() {
		var i Monster
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Basehp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
