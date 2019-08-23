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

package tools

import (
    "log"
)

func Check(e error) {
    if e != nil {
        log.Panic(e)
    }
}
