
-- +migrate Up
CREATE TABLE categories (id integer(11), parent_id integer(11), description varchar(64));

-- +migrate Down
DROP TABLE categories;
