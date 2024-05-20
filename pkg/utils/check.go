package utils

import "github.com/labstack/gommon/log"

func Error(err error) bool {
	if err != nil {
		log.Error(err)
		return true
	}
	return false
}
