
-- +migrate Up
CREATE TABLE IF NOT EXISTS articles (
    id CHAR(36) NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    category_id CHAR(36) NOT NULL,
    status VARCHAR(255) NOT NULL DEFAULT "Draft",
    published_at DATETIME,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- +migrate Down
DROP TABLE IF EXISTS articles;
