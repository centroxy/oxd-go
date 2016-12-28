package utils

import (
	"github.com/juju/loggo"
	"os"
)



func CheckError(prefix string, message string, err error) {
	if err != nil {
		loggo.GetLogger("prefix").Errorf(message,err.Error())
		os.Exit(1)
	}
}
