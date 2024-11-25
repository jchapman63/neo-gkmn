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
    monsterID = @monster_id::uuid
    AND statType = @stat_type::text;

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

-- name: CreateStatForMon :exec
INSERT INTO stats (monsterID, statType, power)
    VALUES ($1, $2, $3);

-- name: RegisterMoveForMon :exec
INSERT INTO movemap (monsterID, moveID)
    VALUES ($1, $2);

