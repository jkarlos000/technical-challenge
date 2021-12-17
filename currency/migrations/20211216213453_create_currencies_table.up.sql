CREATE TABLE IF NOT EXISTS currencies
(
    id
        serial
        PRIMARY
            KEY,
    base        VARCHAR(3) NOT NULL,
    destination VARCHAR(3) NOT NULL,
    rate        DECIMAL(20,
                    6)    DEFAULT 0,
    created_at  TIMESTAMPTZ  NOT NULL,
    updated_at  TIMESTAMPTZ  NOT NULL,
    UNIQUE
        (
         base,
         destination
            )
);