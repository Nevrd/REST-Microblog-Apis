CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT,
    text TEXT NOT NULL,
    tag VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(title)
);
