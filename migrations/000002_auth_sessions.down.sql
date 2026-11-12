DROP INDEX IF EXISTS zapiski.auth_sessions_current_jti_key;
DROP INDEX IF EXISTS zapiski.auth_sessions_user_id_idx;
DROP INDEX IF EXISTS zapiski.auth_sessions_family_id_idx;

DROP TABLE IF EXISTS zapiski.auth_sessions;

DROP TYPE IF EXISTS zapiski.session_status;