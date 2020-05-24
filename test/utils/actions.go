package utils

import (
	"errors"
	"fmt"
)

type actionFunc func(string) (string, error)

var actionFuncs = make(map[string]actionFunc)

func RegisterAction(action string, fn actionFunc) {
	actionFuncs[action] = fn
}

func GetAction(action string) actionFunc {
	return actionFuncs[action]
}

func RunAction(action string, paramFile string) (string, error) {
	fn := GetAction(action)
	if fn != nil {
		return fn(paramFile)
	}
	return "", errors.New(fmt.Sprintf("no registered function for %s", action))
}
