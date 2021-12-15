CREATE TABLE IF NOT EXISTS beers
(
    beer_id
             serial
        PRIMARY
            KEY,
    name
             VARCHAR(50) UNIQUE NOT NULL,
    brewery  VARCHAR(50) UNIQUE NOT NULL,
    country  VARCHAR(70) UNIQUE NOT NULL,
    price    DECIMAL(10, 1) DEFAULT 0,
    currency VARCHAR(3)     DEFAULT 'USD'
);