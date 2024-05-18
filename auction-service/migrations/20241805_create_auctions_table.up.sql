CREATE TABLE auctions (
    id SERIAL PRIMARY KEY,
    item VARCHAR(100) NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
