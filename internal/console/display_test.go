package console_test

import (
	"testing"

	"github.com/sanderdescamps/mgit/internal/console"
)

func TestOutput(t *testing.T) {
	o := console.Display{}

	o.Error("error test message")
	o.Warning("warning test message")
	o.Info("info test message")
	o.Change("change test message")
	o.Ok("ok test message")
	o.Skip("skip test message")
}
