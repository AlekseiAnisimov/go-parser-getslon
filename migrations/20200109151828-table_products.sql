
-- +migrate Up
CREATE TABLE products (
    product_id varchar(64) unique,
    article varchar(64),
    name varchar(255),
    description text,
    available varchar(1) not null,
    merchant_id integer,
    gs_product_key varchar(64),
    gs_category_id integer,
    picture varchar(255),
    thumbnail varchar(255),
    original_picture text,
    vendor varchar(255),
    model varchar(255),
    oldprice float,
    url varchar(254),
    destination_url text,
    currency_id varchar(16),
    price float,
    age varchar(64),
    composition varchar(64),
    other_pictures JSON
);

-- +migrate Down
DROP TABLE products;