-- db/migration/1721225409_create_users.sql

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    password_hash VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    is_password_expired BOOLEAN DEFAULT false
);