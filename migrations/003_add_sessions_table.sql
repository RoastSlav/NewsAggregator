CREATE TABLE sessions
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    session_token TEXT      NOT NULL,
    user_id       INT       NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);