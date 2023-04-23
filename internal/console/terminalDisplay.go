package console

import (
	"os"
)

type TerminalDisplay struct {
	baseDisplay
}

func NewTerminalDisplay(level LogLevel) *TerminalDisplay {
	return &TerminalDisplay{
		baseDisplay{
			LogLevel:  level,
			writerErr: os.Stderr,
			writerOut: os.Stdout,
		},
	}
}

func (d *TerminalDisplay) GetSubDisplay() Display {
	return &TerminalDisplay{
		baseDisplay{
			LogLevel:  d.LogLevel,
			writerErr: os.Stderr,
			writerOut: os.Stdout,
			indent:    d.indent + 2,
		},
	}
}
