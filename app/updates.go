/*
Copyright (c) 2019, Aniol Marti
This file is part of Nininini Bot.
Nininini Bot is free software: you can redistribute it and/or modify
it under the terms of the BSD 3 clause license.
Nininini Bot is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
BSD 3 clause license for more details.
You should have received a copy of the BSD 3 clause license
along with Nininini Bot. If not, see <https://opensource.org/licenses/BSD-3-Clause>.
*/

package app

import (
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/aniolm9/ninini-bot/tools"
)

func ProcessUpdates(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
    for update := range updates {
        if update.InlineQuery == nil {
            continue
        } else if update.InlineQuery.Query == "" {
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
