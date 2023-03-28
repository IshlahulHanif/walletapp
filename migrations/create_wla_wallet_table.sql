CREATE TABLE IF NOT EXISTS wla_wallet
(
    id          UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    customer_id UUID                     NOT NULL,
    is_enabled  BOOLEAN                  NOT NULL DEFAULT false,
    balance     NUMERIC(10, 2)           NOT NULL DEFAULT 0,
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by  VARCHAR(255)             NOT NULL DEFAULT '',
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  VARCHAR(255)             NOT NULL DEFAULT '',
    FOREIGN KEY (customer_id) REFERENCES wla_user_provider (customer_id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_wallets_user_id ON wla_wallet (customer_id);