package wjLibTBot
import innerTelegramBotter "github.com/ancestortelegram/wjLibTBot/internal"
//import innerTelegramBotter "github.com/ancestortelegram/wjLibTBot/internel"

func WJLTB_NewBotter(id int) (IWJTelegramBotAPI, error) {
	return innerTelegramBotter.newTelegramBotApi(id)
}
