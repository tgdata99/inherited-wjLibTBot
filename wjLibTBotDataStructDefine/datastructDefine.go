package wjLibTBotDataStructDefine

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type IWJTelegramBotSPI interface {
	LoadingMeInfo(*SWJTBUser)
	OnError(string, error)
	OnWarning(string)
	BotUpdateFirst(tgbotapi.Update) bool      // don't process this message anymore if return true
	MessageUserCommand(int64, string, string) // userid, command, message;
	MessageUserPhoto(int64, []tgbotapi.PhotoSize)
	MessageUserAudio(int64, *tgbotapi.Audio)
	MessageUserVideo(int64, *tgbotapi.Video)
	MessageUserText(int64, string)
	MessageGroupText(int64, int64, string) // groupid, senduserid, text;
	EventGroupNewJoin(int64, int64)        // groupid ,userid;
	EventGroupLeaveMemb(int64, int64)      // groupid, userid;
}

// interface wises jumper telegram botter;
type IWJTelegramBotAPI interface {
	Initialize(botToken string, spi IWJTelegramBotSPI) error
	Stop()
	GetInfoMe() SWJTBUser
	GetChatInfoByID(int64) (SWJTBUser, error)
	SendTextToID(int64, string) (tgbotapi.Message, error)
	SendImageUrlToID(int64, string, string) (tgbotapi.Message, error)
	SendImageUrlListToID(int64, string, []string) ([]tgbotapi.Message, error)
	GetDirectUrlByFileID(fileid string) (string, error)
	SendChattable(msg tgbotapi.Chattable) (tgbotapi.Message, error)
	GroupDeleteMessage(int64, int) error
	GroupBanMember(int64, int64) error
}

// struct wise jumper telegram botter user;
type SWJTBUser struct {
	Userid       int64
	IsBot        bool
	FirstName    string
	LastName     string
	UserName     string
	Bio          string
	PhotoBigID   string
	PhotoSmallID string
}
