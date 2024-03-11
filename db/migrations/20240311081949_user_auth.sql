-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS auth_methods (
	id 			INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name 		VARCHAR(255) NOT NULL
);

INSERT INTO auth_methods (name) VALUES 
	('email'), 
	('google'), 
	('facebook'), 
	('github');

CREATE TABLE IF NOT EXISTS user_auths (
	id 				INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	first_name 		VARCHAR(255),
	last_name 		VARCHAR(255),
	email 			VARCHAR(255) NOT NULL UNIQUE,
	last_auth_at	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_at 		TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at 		TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_auth_methods (
	id 				INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id 		INT REFERENCES user_auths(id),
    auth_method_id	INT REFERENCES auth_methods(id),
	last_auth_at	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

	UNIQUE (user_id, auth_method_id)
);

INSERT INTO user_auths (email, first_name, last_name) VALUES
	('timurmazitov000@gmail.com', 'Timur', 'Mazitov');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_auth_methods;
DROP TABLE IF EXISTS user_auths;
DROP TABLE IF EXISTS auth_methods;
-- +goose StatementEnd
