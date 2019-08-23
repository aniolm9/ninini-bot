package main

import (
    "strings"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/aniolm9/ninini-bot/tools"
)

func main() {
    data, err := ioutil.ReadFile("TOKEN")
    tools.check(err)
    TOKEN := strings.Trim(string(data), "\n")

    bot, err := tgbotapi.NewBotAPI(TOKEN)
    tools.check(err)
}
