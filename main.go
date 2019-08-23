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

package main

import (
    "strings"
    "log"
    "io/ioutil"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/aniolm9/ninini-bot/tools"
    "github.com/aniolm9/ninini-bot/app"
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
    tools.Check(err)

    app.ProcessUpdates(updates, bot)
}
