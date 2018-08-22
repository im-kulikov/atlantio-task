CREATE TABLE transactions (
  hash VARCHAR(128) NOT NULL PRIMARY KEY,
  "from" VARCHAR(64) NOT NULL,
  "to" VARCHAR(64) NOT NULL,
  block_number INT,
  "seen" bool,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS tx_block_number_idx ON transactions (block_number);
CREATE INDEX IF NOT EXISTS tx_to_idx ON transactions ("to");