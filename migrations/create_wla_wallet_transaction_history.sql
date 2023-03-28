CREATE TABLE IF NOT EXISTS wla_wallet_transaction_history
(
    id            UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    status        VARCHAR(255)             NOT NULL DEFAULT 'pending',
    wallet_id     UUID                  NOT NULL,
    transacted_at TIMESTAMP WITH TIME ZONE NOT NULL,
    type          VARCHAR(255)             NOT NULL,
    amount        NUMERIC(10, 2)           NOT NULL DEFAULT 0,
    reference_id  UUID                     NOT NULL UNIQUE,
    FOREIGN KEY (wallet_id) REFERENCES wla_wallet (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_wallet_transaction_history_reference_id ON wla_wallet_transaction_history (reference_id);
CREATE INDEX IF NOT EXISTS idx_wallet_transaction_history_transacted_at ON wla_wallet_transaction_history (transacted_at);
