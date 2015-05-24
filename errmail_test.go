package errmail_test

import (
	"errors"
	"net/smtp"
	"os"
	"testing"

	"github.com/daaku/errmail"
	"github.com/facebookgo/stackerr"
)

func envLogger() *errmail.Logger {
	return &errmail.Logger{
		Addr: os.Getenv("ERRMAIL_ADDR"),
		Auth: smtp.PlainAuth(
			os.Getenv("ERRMAIL_IDENTITY"),
			os.Getenv("ERRMAIL_USERNAME"),
			os.Getenv("ERRMAIL_PASSWORD"),
			os.Getenv("ERRMAIL_HOST"),
		),
		From: os.Getenv("ERRMAIL_FROM"),
		To:   []string{os.Getenv("ERRMAIL_TO")},
	}
}

func TestSendSimple(t *testing.T) {
	envLogger().Log(errors.New("a simple error"))
}

func TestSendWithStack(t *testing.T) {
	envLogger().Log(stackerr.Wrap(errors.New("a stacked error")))
}
