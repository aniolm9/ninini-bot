package main

import (
    "regexp"
)

func normalToNinini(normalText string) string {
    rLower := regexp.MustCompile("[aeouàáèéòóùú]")
    rUpper := regexp.MustCompile("[AEOUÀÁÈÉÒÓÙÚ]")

    return rUpper.ReplaceAllString(rLower.ReplaceAllString(normalText, "i"), "I")
}
