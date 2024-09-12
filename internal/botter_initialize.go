package innerTelegramBotter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (pInst *cTBotApi) createBot() error {
	bot, err := tgbotapi.NewBotAPI(pInst.botToken)
	if err != nil {
		return err
	}
	pInst.tgBot = bot
	pInst.tgUpdateConfig = tgbotapi.NewUpdate(0)
	pInst.tgUpdateConfig.Timeout = 10

	return nil
}
func (pInst *cTBotApi) loadingInfoMe() error {
	me, err := pInst.tgBot.GetMe()
	if err != nil {
		return err
	}
	pInst.infoMe.Userid = me.ID
	pInst.infoMe.FirstName = me.FirstName
	pInst.infoMe.LastName = me.LastName
	pInst.infoMe.UserName = me.UserName
	pInst.infoMe.IsBot = me.IsBot

	pInst.spi.LoadingMeInfo(&pInst.infoMe)

	return nil
}
