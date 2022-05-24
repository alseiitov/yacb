package repo

// cryptocurrencies queries
const (
	queryGetCurrencies = `
		SELECT id,
			   symbol,
			   name
		FROM crypto_currencies`
)

// rates
const (
	queryInsertRate = `
		INSERT INTO crypto_rates
			(crypto_currency_id, price, date)
		VALUES 
			($1, $2, $3)`

	queryGetRateAtDatetime = `
		SELECT cc.name, cr.price, cr.date
		FROM crypto_rates cr
			JOIN crypto_currencies cc ON cr.crypto_currency_id = cc.id
		WHERE cc.symbol = $1
			AND date = $2
		ORDER BY date ASC
		LIMIT 1`
)
