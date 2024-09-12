package innerTelegramBotter

import (
	"github.com/ancestortelegram/wjLibTBot/wjLibTBotDataStructDefine"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type cTBotApi struct {
	systemid       int
	spi            wjLibTBotDataStructDefine.IWJTelegramBotSPI
	infoMe         wjLibTBotDataStructDefine.SWJTBUser
	tgBot          *tgbotapi.BotAPI
	botToken       string
	isRunning      bool
	tgUpdateConfig tgbotapi.UpdateConfig
}

func NewTelegramBotApi(id int) (*cTBotApi, error) {
	bot := &cTBotApi{systemid: id, spi: newEmptyTBotSpi()}
	return bot, nil
}

func (pInst *cTBotApi) Initialize(botToken string, spi wjLibTBotDataStructDefine.IWJTelegramBotSPI) error {
	pInst.botToken = botToken
	pInst.spi = spi
	err := pInst.createBot()
	if err != nil {
		return err
	}
	err = pInst.loadingInfoMe()
	if err != nil {
		return err
	}
	// init command;!!!

	pInst.isRunning = true

	go pInst.running()

	return nil
}
func (pInst *cTBotApi) Stop() {
	pInst.isRunning = false
}
func (pInst *cTBotApi) GetInfoMe() wjLibTBotDataStructDefine.SWJTBUser {
	return pInst.infoMe
}
func (pInst *cTBotApi) GetChatInfoByID(chatid int64) (wjLibTBotDataStructDefine.SWJTBUser, error) {
	info, err := pInst.tgBot.GetChat(tgbotapi.ChatInfoConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: chatid}})
	var userinfo1 wjLibTBotDataStructDefine.SWJTBUser
	if err != nil {
		return userinfo1, err
	}
	userinfo1.Userid = chatid
	userinfo1.FirstName = info.FirstName
	userinfo1.LastName = info.LastName
	userinfo1.Bio = info.Bio
	userinfo1.IsBot = false
	if info.Photo != nil {
		userinfo1.PhotoBigID = info.Photo.BigFileID
		userinfo1.PhotoSmallID = info.Photo.SmallFileID
	}

	return userinfo1, nil
}
func (pInst *cTBotApi) GetDirectUrlByFileID(fileid string) (string, error) {
	return pInst.tgBot.GetFileDirectURL(fileid)
}
func (pInst *cTBotApi) SendChattable(msg tgbotapi.Chattable) (tgbotapi.Message, error) {
	return pInst.tgBot.Send(msg)
}
