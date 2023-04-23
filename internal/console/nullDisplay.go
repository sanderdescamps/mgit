package console

import (
	"io"
)

type NullDisplay struct {
	baseDisplay
}

func NewNullDisplay() Display {
	return &NullDisplay{
		baseDisplay{
			LogLevel:  LogLevel(INFO),
			writerErr: io.Discard,
			writerOut: io.Discard,
		},
	}
}
