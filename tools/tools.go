package tools

import (
    "log"
)

func check(e error) {
    if e != nil {
        log.Panic(e)
    }
}
