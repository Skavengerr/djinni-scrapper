package telegram

import (
	"log"
	"time"

	configs "github.com/Skavengerrr/job-scrapper/configs"
	"github.com/Skavengerrr/job-scrapper/internal/scrapper"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	cfg            *configs.Config
	chatID         int64
	scrapingActive bool
}

const (
	commandStart = "start"
	start        = "Hi, choose what you need to find"
	findByUrl    = "/Find_by_url"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(findByUrl),
	),
)

func InitBot(configs *configs.Config) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(configs.TelegramApiKey)
	if err != nil {
		log.Panic("Something went wrong with bot initializing", err)
	}

	bot.Debug = true

	return bot
}

func NewBot(bot *tgbotapi.BotAPI, cfg *configs.Config) *Bot {
	return &Bot{
		bot: bot,
		cfg: cfg,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		b.chatID = update.Message.Chat.ID
		break
	}

	if b.chatID != 0 {
		for {
			go scrapper.StartScraping(b.chatID, b.bot)
			time.Sleep(30 * time.Minute)
		}
	}

	return nil
}
