package console

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

//-----------------------------------------

type Status int64

const (
	NO_STATUS Status = iota
	OK
	CHANGED
	FAILED
	SKIPPED
)

//-----------------------------------------

type LogLevel int64

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)
