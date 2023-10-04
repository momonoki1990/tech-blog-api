
-- +migrate Up
CREATE TABLE IF NOT EXISTS categories(
    id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    display_order int DEFAULT 99,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS categories;
