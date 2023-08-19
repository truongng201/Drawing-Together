CREATE TABLE IF NOT EXISTS "user" (
    "id"            UUID PRIMARY KEY,
    "username"      TEXT NOT NULL,
    "email"         TEXT NOT NULL UNIQUE,
    "password_hash" TEXT NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"    TIMESTAMP NULL DEFAULT NULL,
    "avatar_url"    TEXT NULL DEFAULT NULL,
    
);