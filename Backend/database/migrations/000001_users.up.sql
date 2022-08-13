CREATE TABLE IF NOT EXISTS users (
    id BIGINT unsigned NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) DEFAULT NULL,
    email VARCHAR(255) DEFAULT NULL,
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP default now(),
    PRIMARY KEY (id),
    UNIQUE KEY users_email_unique (email)
)
