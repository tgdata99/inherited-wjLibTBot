package innerTelegramBotter

import (
	"github.com/ancestortelegram/wjLibTBot/wjLibTBotDataStructDefine"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type emptyTBotSpi struct {
}

func newEmptyTBotSpi() wjLibTBotDataStructDefine.IWJTelegramBotSPI {
	return &emptyTBotSpi{}
}
func (pInst *emptyTBotSpi) LoadingMeInfo(*wjLibTBotDataStructDefine.SWJTBUser) {
}
func (pInst *emptyTBotSpi) OnError(string, error) {
}
func (pInst *emptyTBotSpi) OnWarning(string) {
}
func (pInst *emptyTBotSpi) BotUpdateFirst(tgbotapi.Update) bool {
	return false
}
func (pInst *emptyTBotSpi) MessageUserCommand(int64, string, string) {
}
func (pInst *emptyTBotSpi) MessageUserPhoto(int64, []tgbotapi.PhotoSize) {
}
func (pInst *emptyTBotSpi) MessageUserAudio(int64, *tgbotapi.Audio) {
}
func (pInst *emptyTBotSpi) MessageUserVideo(int64, *tgbotapi.Video) {
}
func (pInst *emptyTBotSpi) MessageUserText(int64, string) {
}
