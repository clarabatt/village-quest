-- +goose Up
CREATE TABLE IF NOT EXISTS game(
    id TEXT PRIMARY KEY,
    number INTEGER NOT NULL,
    max_days_played INTEGER NOT NULL,
    players_name TEXT NOT NULL
);

-- +goose Down
DROP TABLE game;
