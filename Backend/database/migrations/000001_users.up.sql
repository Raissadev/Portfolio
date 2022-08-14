CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL NOT NULL primary key,
    name VARCHAR(255) DEFAULT NULL,
    email VARCHAR(255) DEFAULT NULL,
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP default now()
);