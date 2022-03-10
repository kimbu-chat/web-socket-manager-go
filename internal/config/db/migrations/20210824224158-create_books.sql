
-- +migrate Up
CREATE TABLE books(
	id SERIAL PRIMARY KEY,
	name VARCHAR(20),
	description VARCHAR(100)
);

-- +migrate Down
DROP TABLE books;
