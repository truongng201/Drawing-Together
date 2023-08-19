----------------
-- TABLE USER --
----------------

CREATE TABLE IF NOT EXISTS "users" (
    "id"            BIGSERIAL PRIMARY KEY,
    "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"    TIMESTAMP NULL DEFAULT NULL,

    "username"      VARCHAR NOT NULL,
    "email"         VARCHAR NULL UNIQUE,
    "hash_password" TEXT NULL DEFAULT NULL,
    "avatar_url"    TEXT NULL DEFAULT NULL,
    "oauth_id"      TEXT NULL DEFAULT NULL,
    "oauth_type"    TEXT NULL DEFAULT NULL
);

----------------
-- TABLE ROOM --
----------------

CREATE TABLE IF NOT EXISTS "rooms" (
    "id"          BIGSERIAL NOT NULL PRIMARY KEY,
    "created_at"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at"  TIMESTAMP NULL DEFAULT NULL,

    "room_id"     TEXT NOT NULL UNIQUE,
    "owner_id"    INT NOT NULL, 
    "is_private"  BOOLEAN NOT NULL DEFAULT FALSE,
    "expired_at"  TIMESTAMP NULL DEFAULT NULL,
    "max_players" INT NOT NULL DEFAULT 5
);

---------------------
-- TABLES RELATION --
---------------------

ALTER TABLE "rooms" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

-----------------
-- TABLE INDEX --
-----------------

CREATE INDEX ON "rooms" ("room_id");

CREATE INDEX ON "users" ("email");
