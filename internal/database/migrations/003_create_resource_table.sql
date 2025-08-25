-- +goose Up
CREATE TABLE IF NOT EXISTS resource (
    id TEXT NOT NULL PRIMARY KEY,
    turn_id TEXT NOT NULL,
    stone INTEGER NOT NULL DEFAULT 0,
    gold INTEGER NOT NULL DEFAULT 0,
    wood INTEGER NOT NULL DEFAULT 0,
    food INTEGER NOT NULL DEFAULT 0,
    worker INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (turn_id) REFERENCES turn(id)
);

CREATE INDEX IF NOT EXISTS idx_resource_turn_id ON resource(turn_id);
CREATE INDEX IF NOT EXISTS idx_resource_created_at ON resource(created_at);
CREATE INDEX IF NOT EXISTS idx_resource_deleted_at ON resource(deleted_at);

-- +goose Down
DROP INDEX IF EXISTS idx_resource_deleted_at;
DROP INDEX IF EXISTS idx_resource_created_at;
DROP INDEX IF EXISTS idx_resource_turn_id;
DROP TABLE IF EXISTS resource;