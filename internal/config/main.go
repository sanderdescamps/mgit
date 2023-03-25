package config

import "github.com/sanderdescamps/mgit/internal/console"

var display *console.Display

func SetDisplay(d *console.Display) {
	display = d
}
