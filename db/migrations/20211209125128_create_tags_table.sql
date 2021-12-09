-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE tags (
  id int not null auto_increment,
  name varchar(50) not null,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE tags;