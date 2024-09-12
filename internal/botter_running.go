package innerTelegramBotter

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (pInst *cTBotApi) running() {
	for {
		if !pInst.isRunning {
			break
		}
		time.Sleep(time.Millisecond * 300)
		pInst.updateMessage_running()
	}
}

func (pInst *cTBotApi) updateMessage_running() {
	updates, err := pInst.tgBot.GetUpdates(pInst.tgUpdateConfig)
	if err != nil {
		pInst.spi.OnError("update message error", err)
		return
	}
	if len(updates) < 1 {
		pInst.spi.OnError("update message empty", err)
		return
	}

	for _, update := range updates {
		if update.UpdateID < pInst.tgUpdateConfig.Offset {
			pInst.spi.OnWarning("message updateid lower than expect")
			continue
		}
		pInst.tgUpdateConfig.Offset = update.UpdateID + 1

		pInst.processUpdate_running(update)
	}

}

func (pInst *cTBotApi) processUpdate_running(update1 tgbotapi.Update) {

	if pInst.spi.BotUpdateFirst(update1) { // return true, means stop deal with this message;
		return
	}

	if update1.Message == nil { // don't do this now;
		return
	}
	fmt.Println("media groupid: ", update1.Message.MediaGroupID)

	isGroup := update1.Message.Chat.Type == "group" || update1.Message.Chat.Type == "supergroup"
	if isGroup {
		if update1.Message.IsCommand() {
			pInst.processMessageGroupCommand_running(update1.Message)
		} else {
			pInst.processMessageGroup_running(update1.Message)
		}
		pInst.processMessageGroup_running(update1.Message)
	} else {
		if update1.Message.IsCommand() {
			pInst.processMessageUserCommand_running(update1.Message)
		} else {
			pInst.processMessageUser_running(update1.Message)
		}
	}

}
func (pInst *cTBotApi) processMessageGroup_running(tMsg *tgbotapi.Message) {
	// nothing now
}
func (pInst *cTBotApi) processMessageGroupCommand_running(tMsg *tgbotapi.Message) {
	// nothing now
}
func (pInst *cTBotApi) processMessageUser_running(tMsg *tgbotapi.Message) {
	fmt.Println("caption: "+tMsg.Caption+" text: ", tMsg.Text)
	if tMsg.Photo != nil {
		pInst.spi.MessageUserPhoto(tMsg.Chat.ID, tMsg.Photo)
	} else if tMsg.Audio != nil {
		pInst.spi.MessageUserAudio(tMsg.Chat.ID, tMsg.Audio)
	} else if tMsg.Video != nil {
		pInst.spi.MessageUserVideo(tMsg.Chat.ID, tMsg.Video)
	}
	if tMsg.Text != "" {
		pInst.spi.MessageUserText(tMsg.Chat.ID, tMsg.Text)
	}

}
func (pInst *cTBotApi) processMessageUserCommand_running(tMsg *tgbotapi.Message) {

	pInst.spi.MessageUserCommand(tMsg.Chat.ID, tMsg.Command(), tMsg.Text)
}
