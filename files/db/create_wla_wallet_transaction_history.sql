CREATE TABLE IF NOT EXISTS wla_wallet_transaction_history
(
    id           SERIAL PRIMARY KEY,
    reference_id VARCHAR(255)   NOT NULL,
    wallet_id    INTEGER        NOT NULL,
    balance      NUMERIC(10, 2) NOT NULL DEFAULT 0,
    create_time  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by   VARCHAR(255),
    update_time  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by   VARCHAR(255),
    FOREIGN KEY (wallet_id) REFERENCES wla_wallet (id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_wallets_transaction_history_reference_id ON wla_wallet_transaction_history (reference_id);
CREATE INDEX IF NOT EXISTS idx_wla_wallet_transaction_history_create_time ON wla_wallet_transaction_history (create_time);
