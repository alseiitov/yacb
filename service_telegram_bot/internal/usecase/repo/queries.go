package repo

// users queries
const (
	queryUserExist = `
		SELECT EXISTS(
			SELECT * FROM users WHERE id = $1
		)`

	queryInsertUser = `
		INSERT INTO users
			(id, password)
		VALUES ($1, $2)
		RETURNING id`

	queryDeleteUser = `
		DELETE FROM users
		WHERE id = $1`

	queryGetUser = `
		SELECT id, password
		FROM users 
		WHERE id = $1`
)

// subscriptions queries
const (
	queryInsertSubscription = `
		INSERT INTO user_subscriptions
			(user_id, crypto_currency_symbol, update_interval)
		VALUES ($1, $2, $3)
		RETURNING id`

	queryGetAllSubscriptions = `
		SELECT id, user_id, crypto_currency_symbol, update_interval
		FROM user_subscriptions s
		ORDER BY update_interval`

	queryGetUsersSubscriptions = `
		SELECT id, crypto_currency_symbol, update_interval
		FROM user_subscriptions s
		WHERE user_id =$1`

	queryDeleteSubscription = `
		DELETE
		FROM user_subscriptions
		WHERE id = $1`
)
