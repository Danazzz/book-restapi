CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR,
    image_url VARCHAR,
    release_year INTEGER CHECK (release_year BETWEEN 1980 AND 2024),
    price INTEGER NOT NULL,
    total_page INTEGER NOT NULL,
    thickness VARCHAR,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);

-- Inisialisasi Admin
INSERT INTO users (username, password, created_by) 
VALUES 
    ('admin', '$2a$10$7qOtbBo5K3i5DMBkzF7mUOp2.nDd9r5t3OGXZ1/FjfJmCZ2h3SEdW', 'system'); 
-- Password: "password123" di-hash dengan bcrypt