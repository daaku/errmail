package errmail_test

import (
	"errors"
	"testing"

	"github.com/daaku/errmail"
	"github.com/facebookgo/stackerr"
)

func TestSendSimple(t *testing.T) {
	errmail.Send(errors.New("a simple error"))
}

func TestSendWithStack(t *testing.T) {
	errmail.Send(stackerr.Wrap(errors.New("a stacked error")))
}
