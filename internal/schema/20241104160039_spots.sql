-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS spots (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  latitude INTEGER NOT NULL,
  longitude INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS spots;
-- +goose StatementEnd
