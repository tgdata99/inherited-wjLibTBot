package WJLibTBot

func WJLTB_NewBotter(id int) (IWJTelegramBotAPI, error) {
	return newTelegramBotApi(id)
}