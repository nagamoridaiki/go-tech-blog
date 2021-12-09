-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE articles_tags (
  article_id int not null,
  tag_id int not null,
  PRIMARY KEY(article_id, tag_id),
  FOREIGN KEY(article_id) REFERENCES articles(id),
  FOREIGN KEY(tag_id) REFERENCES tags(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE articles_tags;
