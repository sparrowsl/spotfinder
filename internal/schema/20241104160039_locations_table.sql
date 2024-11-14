-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS locations (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  latitude FLOAT NOT NULL,
  longitude FLOAT NOT NULL,
  address TEXT NOT NULL,
  category VARCHAR(100),
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS locations;
-- +goose StatementEnd
