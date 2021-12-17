CREATE TABLE IF NOT EXISTS currencies
(
    base        VARCHAR(3) NOT NULL,
    destination VARCHAR(3) NOT NULL,
    rate        DECIMAL(10,
                    1)    DEFAULT 0,
    created_at  TIMESTAMP  NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    UNIQUE
        (
         base,
         destination
            )
);