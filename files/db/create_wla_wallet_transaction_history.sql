CREATE TABLE IF NOT EXISTS wla_wallet_transaction_history
(
    id           SERIAL PRIMARY KEY, --TODO: think if I should delete this bcs we alr get other PK
    reference_id VARCHAR(255) PRIMARY KEY,
    wallet_id    INTEGER        NOT NULL,
    balance      NUMERIC(10, 2) NOT NULL DEFAULT 0,
    create_time  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by   VARCHAR(255),
    update_time  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by   VARCHAR(255),
    FOREIGN KEY (wallet_id) REFERENCES wla_wallet (id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_wallets_user_id ON wla_wallet_transaction_history (wallet_id);