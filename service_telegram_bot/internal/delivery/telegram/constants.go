package telegram_handler

const (
	cmdStart        = `start`
	cmdCurrentRates = `current_rates`
	cmdSubscribe    = `subscribe`
	cmdUnsubscribe  = `unsubscribe`
	cmdMyUserInfo   = `my_user_info`
)

const (
	msgChooseCmdToStart            = `Hello! Choose command to start!`
	msgChooseCurrencyToGetRate     = `Choose currency!`
	msgChooseCurrencyToSubscribe   = `Choose currency to subscribe!`
	msgSubscribed                  = `Done! Now you will receive %s (%s) rate changes every %s`
	msgChooseCurrencyToUnsubscribe = `Choose currency to unsubscribe!`
	msgChooseIntervalToSubscribe   = `Choose interval!`
	msgUserInfo                    = "ID: `%d`\nPassword: `%s`"
	msgYouHaveNoSubscriptions      = "You have no subscriptions"
	msgUnsubscribed                = "Unsubscribed"
)
