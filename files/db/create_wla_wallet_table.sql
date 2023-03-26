CREATE TABLE IF NOT EXISTS wla_wallet
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER        NOT NULL UNIQUE,
    balance     NUMERIC(10, 2) NOT NULL DEFAULT 0,
    create_time TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by  VARCHAR(255),
    update_time TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES wla_user_provider (id) ON DELETE CASCADE
);
