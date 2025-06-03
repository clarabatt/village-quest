-- +goose Up
CREATE TABLE IF NOT EXISTS turn_history (
    id TEXT NOT NULL PRIMARY KEY,
    game_id TEXT NOT NULL,
    number INTEGER NOT NULL,
    players_action TEXT,
    random_events TEXT,
    initial_resources TEXT,
    final_resources TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (game_id) REFERENCES game(id)
);

CREATE INDEX IF NOT EXISTS idx_turn_history_game_id ON turn_history(game_id);
CREATE INDEX IF NOT EXISTS idx_turn_history_number ON turn_history(game_id, number);

-- +goose Down
DROP INDEX IF EXISTS idx_turn_history_number;
DROP INDEX IF EXISTS idx_turn_history_game_id;
DROP TABLE IF EXISTS turn_history;