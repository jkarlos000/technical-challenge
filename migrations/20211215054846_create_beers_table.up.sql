CREATE TABLE IF NOT EXISTS beers
(
    beer_id
             serial
        PRIMARY
            KEY,
    name
             VARCHAR(50) NOT NULL,
    brewery  VARCHAR(50) NOT NULL,
    country  VARCHAR(70) NOT NULL,
    price    DECIMAL(10, 1) DEFAULT 0,
    currency VARCHAR(3)     DEFAULT 'USD',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    UNIQUE (name, brewery, country, currency)
);