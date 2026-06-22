DO
$$
    BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
            CREATE DOMAIN UINT256 AS NUMERIC
                CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
    ELSE
            ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
            ALTER DOMAIN UINT256 ADD
                CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS addresses (
    guid  VARCHAR PRIMARY KEY,
    address VARCHAR UNIQUE NOT NULL,
    address_type VARCHAR(10)  NOT NULL DEFAULT 'user',
    public_key VARCHAR UNIQUE NOT NULL,
    created_at   TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_address_type CHECK(address_type IN ('user', 'hot', 'cold'))
);
CREATE INDEX IF NOT EXISTS idx_addresses_address ON addresses(address);
CREATE INDEX IF NOT EXISTS idx_address_type ON addresses(address_type);


CREATE TABLE IF NOT EXISTS tokens (
    guid              VARCHAR PRIMARY KEY,
    token_address     VARCHAR UNIQUE NOT NULL,
    decimals          SMALLINT NOT NULL DEFAULT 18,
    token_name        VARCHAR NOT NULL,
    token_symbol      VARCHAR NOT NULL,
    collect_amount    UINT256 NOT NULL,
    cold_amount       UINT256 NOT NULL,
    hot_alert_amount  UINT256 NOT NULL,
    created_at        TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_tokens_token_address ON tokens(token_address);


CREATE TABLE IF NOT EXISTS balances (
    guid              VARCHAR PRIMARY KEY,
    address           VARCHAR UNIQUE NOT NULL,
    token_address     VARCHAR UNIQUE NOT NULL,
    balance           UINT256 NOT NULL DEFAULT 0 CHECK(balance >= 0),
    created_at        TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_balances_address ON balances(address);
CREATE INDEX IF NOT EXISTS idx_balances_token_address ON balances(token_address);


