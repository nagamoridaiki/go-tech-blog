-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE articles
  ADD COLUMN writer_id int null,
  ADD FOREIGN KEY (writer_id) REFERENCES writers(id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE articles
DROP FOREIGN KEY articles_ibfk_1,
DROP COLUMN writer_id;