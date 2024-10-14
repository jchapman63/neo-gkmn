INSERT INTO monster (id, name, type, baseHp)
    VALUES ('00000000-0000-0000-0000-000000000000', 'bulbasaur', 'grass', 33);

INSERT INTO monster (id, name, type, baseHp)
    VALUES ('11111111-1111-1111-1111-111111111111', 'charmander', 'fire', 36);

INSERT INTO stats (monsterID, statType, power)
    VALUES ('00000000-0000-0000-0000-000000000000', 'speed', 34);

INSERT INTO stats (monsterID, statType, power)
    VALUES ('11111111-1111-1111-1111-111111111111', 'speed', 36);

INSERT INTO MOVE (id, name, power, type)
    VALUES ('11111111-1111-1111-1111-111111111111', 'tackle', 10, 'normal');

INSERT INTO movemap (monsterID, moveID)
    VALUES ('00000000-0000-0000-0000-000000000000', '11111111-1111-1111-1111-111111111111');

INSERT INTO movemap (monsterID, moveID)
    VALUES ('11111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111');

