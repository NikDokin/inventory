-- +goose Up
-- +goose StatementBegin
-- Create enums
CREATE TYPE commodity_category AS ENUM('culinary', 'medicinal', 'exotic');
CREATE TYPE transaction_type AS ENUM('sale', 'supply');
-- Create commodities table
CREATE TABLE commodities (
   id VARCHAR(36) PRIMARY KEY,
   "name" VARCHAR(255) NOT NULL,
   sku VARCHAR(50) NOT NULL UNIQUE,
   description TEXT,
   category commodity_category NOT NULL,
   quantity INTEGER NOT NULL CHECK (quantity >= 0),
   package VARCHAR(255) NOT NULL,
   price DECIMAL(10, 2) NOT NULL CHECK (price >= 0)
);
-- Create transactions table
CREATE TABLE transactions (
   id VARCHAR(36) PRIMARY KEY,
   commodity_id VARCHAR(36) NOT NULL REFERENCES commodities(id),
   amount INTEGER NOT NULL,
   "type" transaction_type NOT NULL,
   created_at TIMESTAMP NOT NULL,
   saved_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   note TEXT
);
-- Create index on commonly queried fields
CREATE INDEX idx_transactions_created_at ON transactions(created_at);
CREATE EXTENSION pg_trgm;
CREATE INDEX idx_commodities_name_trigram ON commodities USING gin (name gin_trgm_ops);
-- +goose StatementEnd
--
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_commodities_name_trigram;
DROP INDEX IF EXISTS idx_transactions_created_at;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS commodities;
DROP TYPE IF EXISTS transaction_type;
DROP TYPE IF EXISTS commodity_category;
DROP EXTENSION IF EXISTS pg_trgm;
-- +goose StatementEnd