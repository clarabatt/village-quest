-- +goose Up
CREATE TABLE IF NOT EXISTS game (
    id TEXT PRIMARY KEY,
    number INTEGER NOT NULL,
    turns_played INTEGER NOT NULL DEFAULT 0,
    player_name TEXT NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    is_over BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_game_number ON game(number);
CREATE INDEX IF NOT EXISTS idx_game_player_name ON game(player_name);
CREATE INDEX IF NOT EXISTS idx_game_created_at ON game(created_at);

-- +goose Down
DROP INDEX IF EXISTS idx_game_created_at;
DROP INDEX IF EXISTS idx_game_player_name;
DROP INDEX IF EXISTS idx_game_number;
DROP TABLE IF EXISTS game;