CREATE TABLE IF NOT EXISTS wla_user_provider
(
    customer_id VARCHAR(255) PRIMARY KEY, --TODO: think if I should make this one not primary, since 1 might have many token
    token       VARCHAR(255) NOT NULL UNIQUE,
    create_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by  VARCHAR(255),
    update_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  VARCHAR(255)
);
CREATE INDEX idx_wla_user_provider_token ON wla_user_provider (token);