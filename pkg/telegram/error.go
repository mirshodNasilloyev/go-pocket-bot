package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errorInvalidLink        = errors.New("invalid link")
	errorUnautorized        = errors.New("user is not autorized")
	errorUnabletoSave       = errors.New("unable to save")
	errorUnavailableCommand = errors.New("unavailable command")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Error was found")
	switch err {
	case errorInvalidLink:
		msg.Text = "Bu message link emas"
		b.bot.Send(msg)
	case errorUnautorized:
		msg.Text = "Siz avtorizatsiyadan o'tmagansiz. Buni amalga oshirish uchun /start ni bosing"
		b.bot.Send(msg)
	case errorUnabletoSave:
		msg.Text = "Opps, Saqlashni iloji bo'lmadi. Birozdan so'ng qayta urinib ko'ring"
		b.bot.Send(msg)
	case errorUnavailableCommand:
		msg.Text = "Siz mavjud bo'lmagan comandani berdingiz"
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}

}
