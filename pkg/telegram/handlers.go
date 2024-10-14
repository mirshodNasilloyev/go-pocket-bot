package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/url"
)

const (
	commandStart       = "start"
	replyStartTemplate = "Salom! Pocket akkauntingizda havolalarni saqlash uchun, avvalo menga bunga ruxsat berishingiz kerak. Buning uchun quyidagi havolaga o'ting:\n %s"
	replyAutorizedText = "Siz allaqachon avtorizatsiyadan o'tgansiz\n"
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:

		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Link saved successufully")
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errorInvalidLink
	}
	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return errorUnautorized
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		return errorUnabletoSave
	}
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAutorizationProcess(message)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, replyAutorizedText)
	msg.ReplyToMessageID = message.MessageID
	_, err = b.bot.Send(msg)
	return err
}
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	return errorUnavailableCommand
}
