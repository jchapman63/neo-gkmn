-- name: ListMonsters :many
SELECT * FROM monster;

-- name: FetchStat :one
SELECT * FROM stats WHERE monsterID = $1::UUID AND statType = $2::TEXT;
