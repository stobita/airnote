-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE links_tags (
  id INT NOT NULL AUTO_INCREMENT,
  link_id INT NOT NULL,
  tag_id INT NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id),
  UNIQUE uq_links_tags(link_id, tag_id),
  CONSTRAINT fk_link_id
    FOREIGN KEY (link_id)
    REFERENCES links(id),
  CONSTRAINT fk_tag_id
    FOREIGN KEY (tag_id)
    REFERENCES tags(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS links_tags;
