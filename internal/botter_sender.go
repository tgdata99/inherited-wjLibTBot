package innerTelegramBotter

import (
	"errors"
	"io"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (pInst *cTBotApi) SendTextToID(chatid int64, text string) (tgbotapi.Message, error) {
	if text == "" {
		return tgbotapi.Message{}, errors.New("can not send empty text message")
	}
	msg1 := tgbotapi.NewMessage(chatid, text) //"test2 "+strconv.FormatInt(instSelf.glastchatid, 10))
	msg1.ParseMode = "HTML"
	msg1.DisableWebPagePreview = true
	msg1.Text = text

	return pInst.tgBot.Send(msg1)
}
func (pInst *cTBotApi) SendImageUrlToID(chatid int64, text string, imageUrl string) (tgbotapi.Message, error) {
	return pInst.sendImageUrlToUser(chatid, text, imageUrl)
}
func (pInst *cTBotApi) SendImageUrlListToID(chatid int64, text string, imageUrlList []string) ([]tgbotapi.Message, error) {
	var medias []interface{}
	for _, photourl := range imageUrlList {
		imgData, err := downloadPhotomemory(photourl)
		if err != nil {
			return nil, err
		}
		media1 := tgbotapi.NewInputMediaPhoto(tgbotapi.FileBytes{Name: "11", Bytes: imgData})
		//media1 := tgbotapi.NewInputMediaPhoto(tgbotapi.NewPhoto(userid, tgbotapi.FileBytes{Name: "test1", Bytes: data1})) //tgbotapi.FileURL(photo))
		/// ???? should be wrong
		//media1 := tgbotapi.NewPhoto(chatid, tgbotapi.FileBytes{Name: "test1", Bytes: imgData})
		medias = append(medias, media1)
	}
	msg := tgbotapi.MediaGroupConfig{
		ChatID: chatid,
		Media:  medias,
	}
	return pInst.tgBot.SendMediaGroup(msg)
}

func (pInst *cTBotApi) sendImageUrlToUser(userid int64, caption string, photoUrl string) (tgbotapi.Message, error) {
	/*data1, err := downloadPhotomemory(photoUrl)
	if err != nil {
		return tgbotapi.Message{}, err
	}
	msg1 := tgbotapi.NewPhoto(userid, tgbotapi.FileBytes{Name: "test1", Bytes: data1})
	msg1.Caption = caption

	return pInst.tgBot.Send(msg1)*/
	var medias []interface{}
	imgData, err := downloadPhotomemory(photoUrl)
	if err != nil {
		return tgbotapi.Message{}, err
	}
	media1 := tgbotapi.NewInputMediaPhoto(tgbotapi.FileBytes{Name: "11", Bytes: imgData})
	//media1 := tgbotapi.NewInputMediaPhoto(tgbotapi.NewPhoto(userid, tgbotapi.FileBytes{Name: "test1", Bytes: data1})) //tgbotapi.FileURL(photo))
	/// ???? should be wrong
	//media1 := tgbotapi.NewPhoto(chatid, tgbotapi.FileBytes{Name: "test1", Bytes: imgData})
	medias = append(medias, media1)

	msg := tgbotapi.MediaGroupConfig{
		ChatID: userid,
		Media:  medias,
	}
	_, err = pInst.tgBot.SendMediaGroup(msg)
	return tgbotapi.Message{}, err
}

func downloadPhotomemory(url string) ([]byte, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http statuscode error")
	}

	photoData, err := io.ReadAll(resp.Body)
	return photoData, err
}
