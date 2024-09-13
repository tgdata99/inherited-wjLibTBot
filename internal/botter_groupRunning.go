package innerTelegramBotter

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (pInst *cTBotApi) processMessageGroup_running(tMsg *tgbotapi.Message) {
	if tMsg.Photo != nil {
		//pInst.spi.MessageUserPhoto(tMsg.Chat.ID, tMsg.Photo)
	} else if tMsg.Audio != nil {
		//pInst.spi.MessageUserAudio(tMsg.Chat.ID, tMsg.Audio)
	} else if tMsg.Video != nil {
		//pInst.spi.MessageUserVideo(tMsg.Chat.ID, tMsg.Video)
	}
	if tMsg.Text != "" {
		pInst.spi.MessageGroupText(tMsg.Chat.ID, tMsg.From.ID, tMsg.Text)
	} else {
		a := tMsg.NewChatMembers
		if len(a) > 0 {
			fmt.Println("have new members ")
			for iIndex, member := range a {
				fmt.Print(iIndex, member)
			}
		}

		b := tMsg.LeftChatMember
		if b != nil {
			fmt.Println("member left: ", b)
		}
	}
}
func (pInst *cTBotApi) processMessageGroupCommand_running(tMsg *tgbotapi.Message) {

	pInst.spi.MessageUserCommand(tMsg.Chat.ID, tMsg.Command(), tMsg.Text)
}
