
-- +migrate Up
INSERT INTO tech_blog.categories (id, name, display_order) VALUES
    (UUID(), '技術', 1),
    (UUID(), '他', 99);

-- +migrate Down
DELETE FROM tech_blog.categories WHERE name IN ('技術', '他');