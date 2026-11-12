CREATE SCHEMA IF NOT EXISTS zapiski;

CREATE TYPE zapiski.user_state AS ENUM ('unknown', 'active', 'pending_verification', 'blocked');

CREATE TABLE zapiski.users (
    id BIGSERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    username VARCHAR(32) UNIQUE NOT NULL CHECK(char_length(username) BETWEEN 3 AND 32),
    password VARCHAR(72) CHECK(char_length(password) BETWEEN 6 AND 72),
    email VARCHAR(255) UNIQUE NOT NULL CHECK(email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    state zapiski.user_state NOT NULL DEFAULT 'pending_verification',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE zapiski.oauth_accounts (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT NOT NULL
        REFERENCES zapiski.users(id)
        ON DELETE CASCADE,

    provider VARCHAR(32) NOT NULL,
    provider_user_id VARCHAR(255) NOT NULL,

    email VARCHAR(255),
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    UNIQUE (provider, provider_user_id)
);

CREATE INDEX zapiski_oauth_accounts_user_id_idx
    ON zapiski.oauth_accounts (user_id);