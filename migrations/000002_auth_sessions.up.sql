CREATE TYPE zapiski.session_status AS ENUM (
    'active',
    'revoked',
    'compromised'
);

CREATE TABLE zapiski.auth_sessions (
    id UUID PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES zapiski.users(id) ON DELETE CASCADE,

    family_id UUID NOT NULL,
    status zapiski.session_status NOT NULL DEFAULT 'active',

    current_jti UUID NOT NULL,
    previous_jti UUID,

    refresh_token_hash TEXT NOT NULL,
    previous_refresh_token_hash TEXT,

    user_agent TEXT,
    ip INET,

    issued_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_used_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL,

    revoked_at TIMESTAMPTZ,
    revoked_reason TEXT
);

CREATE UNIQUE INDEX auth_sessions_current_jti_key
    ON zapiski.auth_sessions (current_jti);

CREATE INDEX auth_sessions_user_id_idx
    ON zapiski.auth_sessions (user_id);

CREATE INDEX auth_sessions_family_id_idx
    ON zapiski.auth_sessions (family_id);