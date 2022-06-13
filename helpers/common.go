package helpers

import (
	"log"
	"regexp"
	"strings"
)

func RemoveSpecialChar(value string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err.Error())
	}
	value = re.ReplaceAllString(value, " ")
	return strings.TrimSpace(value)
}
