
-- +goose Up
CREATE TABLE IF NOT EXISTS turn_history (
    id UUID NOT NULL PRIMARY KEY,
    game_id UUID NOT NULL REFERENCES games(id),
    number INT NOT NULL,
    players_action JSONB,
    random_events JSONB,
    initial_resources JSONB,
    final_resources JSONB
);

-- +goose Down
DROP TABLE turn_history;