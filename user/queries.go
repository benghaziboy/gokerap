package user

const (
	createUserTable = `
	CREATE TABLE users (
		id         SERIAL PRIMARY KEY,
		email      VARCHAR UNIQUE,
		password   VARCHAR
	)
	`

	createUserTokenTable = `
	CREATE TABLE user_tokens (
		user_id               SERIAL NOT NULL,
		token                 VARCHAR,
		FOREIGN KEY (user_id) REFERENCES users(id),
		PRIMARY KEY (user_id, token)
	)
	`

	createUser = `
	INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id
	`

	createUserToken = `
	INSERT INTO user_tokens (user_id, token) VALUES ($1, $2) RETURNING token
	`

	getUser = `
	SELECT     id, email, password
	FROM       users
	WHERE      id = $1
	`

	getUserByEmail = `
	SELECT     u.id, u.email, u.password, t.token
	FROM       users AS u
	JOIN       user_tokens AS t
	ON         u.id = t.user_id
	WHERE      u.email = $1
	`

	getUserByToken = `
	SELECT     u.id, u.email
	FROM       user_tokens AS t
	JOIN       users AS u
	ON         t.user_id = u.id
	WHERE      t.token = $1
	`
)
