package main

import (
    "strings"
    "log"
    "io/ioutil"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/aniolm9/ninini-bot/tools"
)

func main() {
    data, err := ioutil.ReadFile("TOKEN")
    tools.Check(err)
    TOKEN := strings.Trim(string(data), "\n")

    bot, err := tgbotapi.NewBotAPI(TOKEN)
    tools.Check(err)
    bot.Debug = false
    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.InlineQuery.Query == "" {
            continue
        }
        log.Printf("[%s] %s", update.InlineQuery.From, update.InlineQuery.Query)
        nininiString := normalToNinini(update.InlineQuery.Query)
        article := tgbotapi.NewInlineQueryResultArticle("nininiString", nininiString, nininiString)

        inlineConf := tgbotapi.InlineConfig{
            InlineQueryID:  update.InlineQuery.ID,
            IsPersonal:     true,
            CacheTime:      0,
            Results:        []interface{}{article},
        }
        _, err := bot.AnswerInlineQuery(inlineConf)
        tools.Check(err)
    }
}
