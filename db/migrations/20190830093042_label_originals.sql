-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE link_originals (
  id INT NOT NULL AUTO_INCREMENT,
  link_id INT NOT NULL,
  title text,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id),
  UNIQUE uq_link_originals(link_id),
  CONSTRAINT fk_link_originals_link_id
    FOREIGN KEY (link_id)
    REFERENCES links(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS link_originals;
