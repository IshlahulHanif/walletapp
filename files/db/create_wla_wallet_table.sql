CREATE TABLE IF NOT EXISTS wla_wallet
(
    id          SERIAL PRIMARY KEY,
    customer_id VARCHAR(255)   NOT NULL,
    is_enabled  BOOLEAN        NOT NULL DEFAULT false,
    balance     NUMERIC(10, 2) NOT NULL DEFAULT 0,
    create_time TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by  VARCHAR(255),
    update_time TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  VARCHAR(255),
    FOREIGN KEY (customer_id) REFERENCES wla_user_provider (customer_id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_wallets_user_id ON wla_wallet (customer_id);