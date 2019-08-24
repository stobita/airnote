-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE tags (
  id INT NOT NULL AUTO_INCREMENT,
  text text NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS tags;
