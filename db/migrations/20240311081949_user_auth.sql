-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS auth_methods (
	id 			SERIAL PRIMARY KEY,
	name 		VARCHAR(255) NOT NULL
);

INSERT INTO auth_methods (name) VALUES 
	('email'), 
	('google'), 
	('facebook'), 
	('github');

CREATE TABLE IF NOT EXISTS users_auth (
	id 			SERIAL PRIMARY KEY,
	first_name 	VARCHAR(255) NOT NULL,
	last_name 	VARCHAR(255) NOT NULL,
	email 		VARCHAR(255) NOT NULL UNIQUE,
	created_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users_auth_methods (
    user_id 		INT REFERENCES users_auth(id),
    auth_method_id 	INT REFERENCES auth_methods(id),
    PRIMARY KEY (user_id, auth_method_id)
);

INSERT INTO users_auth (email, first_name, last_name) VALUES
	('timurmazitov000@gmail.com', 'Timur', 'Mazitov');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_auth_methods;
DROP TABLE IF EXISTS users_auth;
DROP TABLE IF EXISTS auth_methods;
-- +goose StatementEnd
