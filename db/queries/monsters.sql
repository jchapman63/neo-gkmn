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

-- name: FetchMovesForMon :many
SELECT
    *
FROM
    movemap
WHERE
    monsterid = $1;

-- name: CreateMonster :exec
INSERT INTO monster (id, name, type, baseHp)
    VALUES ($1, $2, $3, $4);

-- name: CreateMove :exec
INSERT INTO MOVE (id, name, power, type)
    VALUES ($1, $2, $3, $4);

