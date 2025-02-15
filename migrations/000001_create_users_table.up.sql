CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL,
    avatar_url text,
    cover_image_url text,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);
