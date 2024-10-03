CREATE TABLE categories
(
    id   INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

ALTER TABLE articles
    ADD COLUMN category_id INT UNSIGNED,
    ADD CONSTRAINT category_article_fk
        FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE NO ACTION;