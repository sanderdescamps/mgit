package console

import (
	"fmt"
	"io"
	"strings"
)

type baseDisplay struct {
	Display
	writerErr io.Writer
	writerOut io.Writer
	LogLevel  LogLevel
	indent    int
}

func logFormat(prefix string, msg string, color Color, indent int) string {
	return fmt.Sprintf("%s%s%s: %s%s", strings.Repeat(" ", indent), color, strings.ToUpper(prefix), msg, RESET)
}

func statusFormat(prefix string, msg string, color Color, indent int) string {
	return fmt.Sprintf("%s=>%s%s: %s%s", strings.Repeat(" ", indent), color, strings.ToUpper(prefix), msg, RESET)
}

func (d *baseDisplay) Errorf(msg string, a ...any) {
	d.Error(fmt.Sprintf(msg, a...))
}

func (d *baseDisplay) Warningf(msg string, a ...any) {
	d.Warning(fmt.Sprintf(msg, a...))
}

func (d *baseDisplay) Infof(msg string, a ...any) {
	d.Info(fmt.Sprintf(msg, a...))
}

func (d *baseDisplay) Debugf(msg string, a ...any) {
	d.Debug(fmt.Sprintf(msg, a...))
}

func (d *baseDisplay) Error(msg string) {
	d.PrintErr(logFormat("error", msg, Color(RED), d.indent))
}

func (d *baseDisplay) Warning(msg string) {
	d.PrintErr(logFormat("warning", msg, Color(YELLOW), d.indent))
}

func (d *baseDisplay) Info(msg string) {
	d.PrintErr(logFormat("info", msg, Color(WHITE), d.indent))
}

func (d *baseDisplay) Debug(msg string) {
	d.PrintErr(logFormat("debug", msg, Color(WHITE), d.indent))
}

func (d *baseDisplay) Print(msg string) {
	fmt.Fprint(d.writerOut, msg+"\n")
}

func (d *baseDisplay) Printf(msg string, a ...any) {
	d.Print(fmt.Sprintf(msg, a...))
}
func (d *baseDisplay) PrintColorf(msg string, color Color, a ...any) {
	msg = fmt.Sprintf("%s%s%s", color.String(), msg, RESET.String())
	d.Print(fmt.Sprintf(msg, a...))
}

func (d *baseDisplay) PrintErr(msg string) {
	fmt.Fprint(d.writerErr, msg+"\n")
}

func (d *TerminalDisplay) DebugWriter() io.Writer {
	if d.LogLevel >= ERROR {
		return NewTransformWriter(d.writerErr, func(b []byte) []byte {
			return []byte(logFormat("debug", string(b), Color(WHITE), d.indent))
		})
	} else {
		return io.Discard
	}
}

func (d *TerminalDisplay) InfoWriter() io.Writer {
	if d.LogLevel >= INFO {
		return NewTransformWriter(d.writerErr, func(b []byte) []byte {
			return []byte(logFormat("info", string(b), Color(WHITE), d.indent))
		})
	} else {
		return io.Discard
	}
}

func (d *TerminalDisplay) WarningWriter() io.Writer {
	if d.LogLevel >= WARNING {
		return NewTransformWriter(d.writerErr, func(b []byte) []byte {
			return []byte(logFormat("warning", string(b), Color(YELLOW), d.indent))
		})
	} else {
		return io.Discard
	}
}

func (d *TerminalDisplay) ErrorWriter() io.Writer {
	if d.LogLevel >= ERROR {
		return NewTransformWriter(d.writerErr, func(b []byte) []byte {
			return []byte(logFormat("error", string(b), Color(RED), d.indent))
		})
	} else {
		return io.Discard
	}
}

func (d *baseDisplay) Final(status Status, msg string) {
	switch status {
	case NO_STATUS:
		d.Print(statusFormat("no status", msg, Color(PURPLE), d.indent))
	case OK:
		d.Print(statusFormat("ok", msg, Color(GREEN), d.indent))
	case CHANGED:
		d.Print(statusFormat("change", msg, Color(YELLOW), d.indent))
	case SKIPPED:
		d.Print(statusFormat("skip", msg, Color(CYAN), d.indent))
	case FAILED:
		d.Print(statusFormat("failed", msg, Color(RED), d.indent))
	default:
		d.Print(statusFormat("invalid status", msg, Color(PURPLE), d.indent))
	}
}

func (d *baseDisplay) GetSubDisplay() Display {
	return &baseDisplay{
		writerErr: d.writerErr,
		writerOut: d.writerOut,
		LogLevel:  d.LogLevel,
		indent:    d.indent + 2,
	}
}
