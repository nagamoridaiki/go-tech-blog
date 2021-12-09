-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO writers (id, name) VALUES
(1001, 'ライターA'),
(1002, 'ライターB');

INSERT INTO articles (id, title, body, created, updated, writer_id) VALUES
(9001, 'ウェブサイトを作ろう 1', 'ダミーテキスト', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1001),
(9002, 'ウェブサイトを作ろう 2', 'ダミーテキスト', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1001),
(9003, 'ウェブサイトを作ろう 3', 'ダミーテキスト', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1001),
(9004, 'Goでアプリケーションを作ろう 1', 'ダミーテキスト', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1002),
(9005, 'Goでアプリケーションを作ろう 2', 'ダミーテキスト', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 1002);

INSERT INTO tags (id, name) VALUES
(7001, 'HTML'),
(7002, 'CSS'),
(7003, 'JavaScript'),
(7004, 'Go'),
(7005, 'Linux'),
(7006, 'Database');

INSERT INTO articles_tags (article_id, tag_id) VALUES
(9001, 7001),
(9001, 7002),
(9002, 7001),
(9002, 7002),
(9002, 7003),
(9003, 7003),
(9004, 7004),
(9004, 7005),
(9005, 7004),
(9005, 7006);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM articles_tags WHERE article_id IN(9001, 9002, 9003, 9004, 9005);
DELETE FROM articles WHERE id IN(9001, 9002, 9003, 9004, 9005);
DELETE FROM writers WHERE id IN(1001, 1002);
DELETE FROM tags WHERE id IN(7001, 7002, 7003, 7004, 7005, 7006);
