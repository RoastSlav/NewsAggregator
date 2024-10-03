CREATE TABLE IF NOT EXISTS comments
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    article_id INT  NOT NULL,
    user_id    INT  NOT NULL,
    content    TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE comments
    ADD CONSTRAINT comments_article_fk FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
    ADD CONSTRAINT comments_user_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS likes
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    article_id INT NOT NULL,
    user_id    INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_like (article_id, user_id)
);

ALTER TABLE likes
    ADD CONSTRAINT likes_article_fk FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
    ADD CONSTRAINT likes_user_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS read_later
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    article_id INT NOT NULL,
    user_id    INT NOT NULL,
    added_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_read_later (article_id, user_id)
);

ALTER TABLE read_later
    ADD CONSTRAINT read_later_article_fk FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
    ADD CONSTRAINT read_later_user_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;
