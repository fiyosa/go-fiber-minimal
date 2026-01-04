package lang

import (
	"fmt"
	"go-fiber-minimal/config"
	"strings"
)

var Trans transManager

type transManager struct{}

func (transManager) Convert(msg string, args ...map[string]any) string {
	if len(args) == 0 || args[0] == nil {
		return msg
	}

	newMsg := msg
	for key, value := range args[0] {
		newMsg = strings.ReplaceAll(newMsg, ":"+key, fmt.Sprintf("%v", value))
	}
	return newMsg
}

func (transManager) Get() ILang {
	return map[string]ILang{
		"en": EN,
		"id": ID,
	}[config.Env.APP_LOCALE]
}
