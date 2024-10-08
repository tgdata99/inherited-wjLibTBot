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
func (pInst *emptyTBotSpi) MessageUserCommand(*tgbotapi.Message, int64, string, string) {
}
func (pInst *emptyTBotSpi) MessageUserPhoto(*tgbotapi.Message, int64, []tgbotapi.PhotoSize) {
}
func (pInst *emptyTBotSpi) MessageUserAudio(*tgbotapi.Message, int64, *tgbotapi.Audio) {
}
func (pInst *emptyTBotSpi) MessageUserVideo(*tgbotapi.Message, int64, *tgbotapi.Video) {
}
func (pInst *emptyTBotSpi) MessageUserText(*tgbotapi.Message, int64, string) {
}
func (pInst *emptyTBotSpi) MessageGroupText(*tgbotapi.Message, int64, int64, string) {
}
func (pInst *emptyTBotSpi) MessageGroupPhoto(*tgbotapi.Message, int64, int64, []tgbotapi.PhotoSize) {
}
func (pInst *emptyTBotSpi) MessageGroupVideo(*tgbotapi.Message, int64, int64, *tgbotapi.Video) {
}
func (pInst *emptyTBotSpi) EventGroupNewJoin(*tgbotapi.Message, int64, int64) {
}
func (pInst *emptyTBotSpi) EventGroupLeaveMember(*tgbotapi.Message, int64, int64) {
}
func (pInst *emptyTBotSpi) MessageGroupCommand(*tgbotapi.Message, int64, int64, string, string) {
}
