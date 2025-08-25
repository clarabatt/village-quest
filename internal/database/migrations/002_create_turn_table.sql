-- +goose Up
CREATE TABLE IF NOT EXISTS turn (
    id TEXT NOT NULL PRIMARY KEY,
    game_id TEXT NOT NULL,
    number INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (game_id) REFERENCES game(id)
);

CREATE INDEX IF NOT EXISTS idx_turn_game_id ON turn(game_id);
CREATE INDEX IF NOT EXISTS idx_turn_number ON turn(game_id, number);
CREATE INDEX IF NOT EXISTS idx_turn_created_at ON turn(created_at);

-- +goose Down
DROP INDEX IF EXISTS idx_turn_created_at;
DROP INDEX IF EXISTS idx_turn_number;
DROP INDEX IF EXISTS idx_turn_game_id;
DROP TABLE IF EXISTS turn;