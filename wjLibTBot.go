package wjLibTBot

import "github.com/ancestortelegram/wjLibTBot/wjLibTBotDataStructDefine"
import innerTelegramBotter "github.com/ancestortelegram/wjLibTBot/internal"
//import innerTelegramBotter "github.com/ancestortelegram/wjLibTBot/internal"

func WJLTB_NewBotter(id int) (wjLibTBotDataStructDefine.IWJTelegramBotAPI, error) {
	return innerTelegramBotter.NewTelegramBotApi(id)
}
