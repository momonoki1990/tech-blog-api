
-- +migrate Up
CREATE TABLE IF NOT EXISTS taggings (
    article_id CHAR(36) NOT NULL,
    tag_name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (tag_name) REFERENCES tags(name),
    PRIMARY KEY (article_id, tag_name)
);

-- +migrate Down
DROP TABLE IF EXISTS taggings;
