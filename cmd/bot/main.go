package main

import (
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mirshodNasilloyev/tg-bot-youtube-go/pkg/config"
	"github.com/mirshodNasilloyev/tg-bot-youtube-go/pkg/repository"
	"github.com/mirshodNasilloyev/tg-bot-youtube-go/pkg/repository/boltdb"
	"github.com/mirshodNasilloyev/tg-bot-youtube-go/pkg/server"
	"github.com/mirshodNasilloyev/tg-bot-youtube-go/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal("Initializing config error", err)
	}
	log.Println(cfg)

	bot, err := tgbotapi.NewBotAPI("7740158252:AAFoi8HjtQEiP2dvuY8SlqDop2vYn5vjzH8")
	if err != nil {
		log.Fatal("Incorrect token: ", err)
	}
	bot.Debug = true

	pocketClient, err := pocket.NewClient("112458-1d1ec5143833166618cd3bd")
	if err != nil {
		log.Fatal("Incorrect pocket client: ", err)
	}
	db, err := initDB()
	if err != nil {
		log.Fatal("Incorrect db: ", err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)
	authorizedServer := server.NewAuthorizationServer(pocketClient, tokenRepository, "https://t.me/hr_managment_bot")

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")
	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizedServer.Start(); err != nil {
		log.Fatal(err)
	}

}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessToken))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestToken))
		if err != nil {
			return err
		}
		return nil
	})
	return db, nil
}
