
-- +migrate Up
ALTER TABLE users
ADD COLUMN authToken VARCHAR;

-- +migrate Down
ALTER TABLE users
DROP COLUMN authToken;