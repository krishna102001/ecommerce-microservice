CREATE TABLE IF NOT EXISTS orders(
    id CHAR(27) PRIMARY KEY,
    created_at TIMESTAMP WITH ZONE NOT NULL,
    account_id CHAR(27) NOT NULL,
    total_price MONEY NOT NULL
);