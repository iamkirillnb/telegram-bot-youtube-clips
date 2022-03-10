package main

import (
	"flag"
	"telegram_bot/internal"
	"telegram_bot/internal/telebot"
	"telegram_bot/pkg/logging"
)

var cfgPath string

func init(){
	flag.StringVar(&cfgPath, "config", "dev.yaml", "config file path")
}

func main() {
	flag.Parse()

	log := logging.NewLogger()
	cfg := internal.GetConfig(cfgPath)

	bot := telebot.NewBot(log, cfg)
	bot.Run()

}
