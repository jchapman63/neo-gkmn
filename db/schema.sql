CREATE TABLE monster (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    baseHp INTEGER NOT NULL
);

CREATE TABLE move (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    power INTEGER NOT NULL,
    type TEXT NOT NULL
);

CREATE TABLE movemap (
    monsterID UUID REFERENCES monster(id) ON DELETE CASCADE,
    moveID UUID REFERENCES move(id) ON DELETE CASCADE,
    PRIMARY KEY(monsterID, moveID)
);
