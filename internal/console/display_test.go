package console_test

import (
	"testing"

	"github.com/sanderdescamps/mgit/internal/console"
)

func TestOutput(t *testing.T) {
	o := console.NewTerminalDisplay(console.DEBUG)
	o.Error("error test message")
	o.Warning("warning test message")
	o.Info("info test message")
	o.Final(console.CHANGED, "change test message")
	o.Final(console.OK, "ok test message")
	o.Final(console.SKIPPED, "skip test message")
	o.Final(console.FAILED, "skip test message")
}
