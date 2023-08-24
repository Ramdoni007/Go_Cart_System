CREATE TABLE "users"
(
    "username"           varchar PRIMARY KEY,
    "hashed_password"    varchar        NOT NULL,
    "full_name"          varchar        NOT NULL,
    "email"              varchar UNIQUE NOT NULL,
    "password_change_at" timestamptz    NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at"         timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "cart"
(
    "id"           bigserial PRIMARY KEY,
    "name_product" varchar NOT NULL,
    "entity"       bigint  NOT NULL
);