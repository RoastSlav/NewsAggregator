CREATE DATABASe news_aggregator;

USE news_aggregator;

CREATE TABLE articles
(
    id           INT AUTO_INCREMENT PRIMARY KEY,
    source_id    VARCHAR(100),
    source_name  VARCHAR(255),
    author       VARCHAR(255),
    title        VARCHAR(255) NOT NULL,
    description  TEXT,
    url          VARCHAR(255) NOT NULL,
    url_to_image VARCHAR(255),
    published_at DATETIME,
    content      TEXT,
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP
);
