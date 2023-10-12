
-- +migrate Up
CREATE TABLE IF NOT EXISTS taggings (
    article_id CHAR(36) NOT NULL,
    tag_id CHAR(36) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id),
    PRIMARY KEY (article_id, tag_id)
);

-- +migrate Down
DROP TABLE IF EXISTS taggings;
