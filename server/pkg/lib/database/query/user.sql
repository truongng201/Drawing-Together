CREATE TABLE IF NOT EXISTS "user" (
    "id"            UUID PRIMARY KEY,
    "email"         TEXT NOT NULL UNIQUE,
    "password_hash" TEXT NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);