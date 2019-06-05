CREATE TABLE users (
    id         SERIAL,
    email      VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,

    PRIMARY KEY (id),
    UNIQUE (email)
)