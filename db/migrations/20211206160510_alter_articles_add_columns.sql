-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
alter table articles
  add column body mediumtext NOT NULL,
  add column created datetime,
  add column updated datetime;

update articles set created = CURRENT_TIMESTAMP where created is null;
update articles set updated = CURRENT_TIMESTAMP where updated is null;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

alter table articles
  drop column body,
  drop column created,
  drop column updated;
