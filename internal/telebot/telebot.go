package telebot

import (
	"context"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"net/http"
	"telegram_bot/internal"
	"telegram_bot/internal/clients/youtube"
	"telegram_bot/internal/service"
	"telegram_bot/pkg/logging"
	"time"
)

type Bot interface {
	Run()
}

type telegramBot struct {
	cfg            *internal.Config
	log            logging.Logger
	youtubeService service.YoutubeService
}

func NewBot(logger logging.Logger, config *internal.Config) telegramBot {

	ytClient := youtube.NewClientYouTube(config.Youtube.Url, config.Youtube.ApiKey, &http.Client{})
	youtubeService := service.NewYoutubeService(ytClient, &logger)

	return telegramBot{
		cfg:            config,
		log:            logger,
		youtubeService: youtubeService,
	}
}

func (t *telegramBot) Run() {
	t.StartBot()
}

func (t *telegramBot) StartBot() {
	pref := tele.Settings{
		Token:  t.cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		t.log.Fatal(err)
		return
	}

	b.Handle("/m", func(c tele.Context) error {
		treckName := c.Message().Payload

		treck, err := t.youtubeService.FindTrackByName(context.Background(), treckName)
		if err != nil {
			return err
		}
		t.log.Infof("user %s is looking for a song %s", c.Message().Sender.Username, treckName)
		return c.Send(treck)
	})
	b.Handle("/start", func(c tele.Context) error {
		t.log.Infof("user %s start use bot", c.Message().Sender.Username)
		return c.Send(fmt.Sprintf("Привет %s %s", c.Message().Sender.FirstName, c.Message().Sender.LastName))
	})

	b.Start()
}
