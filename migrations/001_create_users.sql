CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(50) NOT NULL CHECK (char_length(first_name) >= 2),
    last_name VARCHAR(50) NOT NULL CHECK (char_length(last_name) >= 2),
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(25),
    age INTEGER CHECK (age > 0),
    status VARCHAR(10) NOT NULL DEFAULT 'Active' CHECK (status IN ('Active', 'Inactive')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
