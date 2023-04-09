package console

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// const Reset = "\033[0m"

// const Red = "\033[31m"
// const Green = "\033[32m"
// const Yellow = "\033[33m"
// const Blue = "\033[34m"
// const Purple = "\033[35m"
// const Cyan = "\033[36m"
// const Gray = "\033[37m"
// const White = "\033[97m"

type Color int64

const (
	RESET Color = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
	CYAN
	GRAY
	ORANGE
	WHITE
)

type LogLevel int64

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

func (c Color) String() string {
	switch c {
	case RESET:
		// return "\x1b[0m"
		return "\033[0m"
	case RED:
		// return "\x1b[31m"
		return "\033[31m"
	case GREEN:
		return "\033[32m"
	case YELLOW:
		return "\033[33m"
	case BLUE:
		return "\033[34m"
	case PURPLE:
		return "\033[35m"
	case CYAN:
		return "\033[36m"
	case GRAY:
		return "\033[90m"
	case WHITE:
		return "\033[97m"
	case ORANGE:
		return "\033[38;5;130m"
		// return "\033[38;5;166m"
	default:
		return "\033[0m"
	}
}

type Display struct {
	LogLevel LogLevel
	writer   io.Writer
}

func NewStdoutDisplay(logLevel LogLevel) *Display {
	return &Display{
		writer: os.Stdout,
	}
}

func (o *Display) SetLogLevel(level LogLevel) {
	o.LogLevel = level
}

func (o Display) Error(msg string) {
	if o.LogLevel <= ERROR {
		o.logPrint("error", msg, RED)
	}
}

func (o Display) Errorf(msg string, a ...any) {
	o.Error(fmt.Sprintf(msg, a...))
}

func (o Display) Warning(msg string) {
	if o.LogLevel <= WARNING {
		o.logPrint("Warning", msg, ORANGE)
	}
}

func (o Display) Warningf(msg string, a ...any) {
	o.Warning(fmt.Sprintf(msg, a...))
}

func (o Display) Info(msg string) {
	if o.LogLevel <= INFO {
		o.logPrint("info", msg, WHITE)
	}
}

func (o Display) Infof(msg string, a ...any) {
	o.Info(fmt.Sprintf(msg, a...))
}

func (o Display) Debug(msg string) {
	if o.LogLevel <= DEBUG {
		o.logPrint("debug", msg, WHITE)
	}
}

func (o Display) Debugf(msg string, a ...any) {
	o.Debug(fmt.Sprintf(msg, a...))
}

func (o Display) Change(msg string) {
	o.logPrint("changed", msg, YELLOW)
}

func (o Display) Changef(msg string, a ...any) {
	o.Change(fmt.Sprintf(msg, a...))
}

func (o Display) Ok(msg string) {
	o.logPrint("ok", msg, GREEN)
}

func (o Display) Okf(msg string, a ...any) {
	o.Ok(fmt.Sprintf(msg, a...))
}

func (o Display) Skip(msg string) {
	o.logPrint("skip", msg, BLUE)
}

func (o Display) Skipf(msg string, a ...any) {
	o.Skip(fmt.Sprintf(msg, a...))
}

func (o Display) Print(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func (o Display) logPrint(prefix string, msg string, color Color) {
	text := fmt.Sprintf("  %s%s: %s%s\n", color, strings.ToUpper(prefix), msg, RESET)
	o.writer.Write([]byte(text))
}

// func (o Display) Writer() *os.File {
// 	return os.Stdout
// }

func (o Display) Writer() *io.Writer {
	return &o.writer
}

func (o Display) InfoWriter() *io.Writer {
	if o.LogLevel <= INFO {
		return NewTransformWriter(o.writer, func(b []byte) []byte {
			s := string(b)
			s = fmt.Sprintf("info: %s", s)
			b = []byte(s)
			return b
		})
	} else {
		return &io.Discard
	}
}

func (o Display) ErrorWriter() *io.Writer {
	if o.LogLevel <= ERROR {
		return &o.writer
	} else {
		return &io.Discard
	}
}

type TransformWriter struct {
	io.Writer
	writer io.Writer
	f      func([]byte) []byte
}

func NewTransformWriter(writer io.Writer, f func([]byte) []byte) *io.Writer {
	var result io.Writer
	result = TransformWriter{writer: writer, f: f}
	return &result
}

func (t TransformWriter) Write(p []byte) (n int, err error) {
	return t.writer.Write(t.f(p))
}
