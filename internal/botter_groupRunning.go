package innerTelegramBotter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (pInst *cTBotApi) processMessageGroup_running(tMsg *tgbotapi.Message) {
	if tMsg.Photo != nil {
		pInst.spi.MessageGroupPhoto(tMsg, tMsg.Chat.ID, tMsg.From.ID, tMsg.Photo)
	} else if tMsg.Audio != nil {
		//pInst.spi.MessageUserAudio(tMsg.Chat.ID, tMsg.Audio)
	} else if tMsg.Video != nil {
		pInst.spi.MessageGroupVideo(tMsg, tMsg.Chat.ID, tMsg.From.ID, tMsg.Video)
	}
	if tMsg.Text != "" {
		pInst.spi.MessageGroupText(tMsg, tMsg.Chat.ID, tMsg.From.ID, tMsg.Text)
	} else {
		a := tMsg.NewChatMembers
		if len(a) > 0 {
			for _, member := range a {
				pInst.spi.EventGroupNewJoin(tMsg, tMsg.Chat.ID, member.ID)
			}
		}

		b := tMsg.LeftChatMember
		if b != nil {
			pInst.spi.EventGroupLeaveMember(tMsg, tMsg.Chat.ID, b.ID)
		}
	}
}
func (pInst *cTBotApi) processMessageGroupCommand_running(tMsg *tgbotapi.Message) {
	pInst.spi.MessageGroupCommand(tMsg, tMsg.Chat.ID, tMsg.From.ID, tMsg.Command(), tMsg.Text)
}
