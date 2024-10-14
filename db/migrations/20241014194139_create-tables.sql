-- migrate:up
CREATE TABLE monster (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name text NOT NULL,
    type TEXT NOT NULL,
    baseHp integer NOT NULL
);

CREATE TABLE MOVE (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name text NOT NULL,
    power integer NOT NULL,
    type TEXT NOT NULL
);

CREATE TABLE movemap (
    monsterID uuid REFERENCES monster (id) ON DELETE CASCADE,
    moveID uuid REFERENCES MOVE (id) ON DELETE CASCADE,
    PRIMARY KEY (monsterID, moveID)
);

CREATE TABLE stats (
    monsterID uuid REFERENCES monster (id) ON DELETE CASCADE,
    statType text NOT NULL,
    power integer NOT NULL,
    PRIMARY KEY (monsterID, statType)
);

-- migrate:down
DROP TABLE monster;

DROP TABLE MOVE;

DROP TABLE movemap;

DROP TABLE stats;

