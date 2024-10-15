-- name: ListMonsters :many
SELECT
    *
FROM
    monster;

-- name: FetchStat :one
SELECT
    *
FROM
    stats
WHERE
    monsterID = $1::uuid
    AND statType = $2::text;

-- name: FetchMonster :one
SELECT
    *
FROM
    monster
WHERE
    id = $1;

-- name: FetchMove :one
SELECT
    *
FROM
    MOVE
WHERE
    id = $1;

