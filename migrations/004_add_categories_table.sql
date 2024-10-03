CREATE TABLE categories
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

ALTER TABLE articles
    ADD COLUMN category_id BIGINT UNSIGNED,
    ADD CONSTRAINT fk_category
        FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE NO ACTION;