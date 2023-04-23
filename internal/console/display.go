package console

import "io"

type Display interface {
	Debug(msg string)
	Debugf(msg string, a ...any)
	Info(msg string)
	Infof(msg string, a ...any)
	Warning(msg string)
	Warningf(msg string, a ...any)
	Error(msg string)
	Errorf(msg string, a ...any)
	Print(msg string)
	Printf(msg string, a ...any)
	PrintColorf(msg string, color Color, a ...any)
	PrintErr(msg string)
	DebugWriter() io.Writer
	InfoWriter() io.Writer
	WarningWriter() io.Writer
	ErrorWriter() io.Writer
	Final(status Status, mst string)
	GetSubDisplay() Display
}
