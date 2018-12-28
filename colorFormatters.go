package main

import (
	"fmt"
	"github.com/aybabtme/color/brush"
)

type ColorFormatter interface {
	direction(s string) string
	place(s string) string
	time(s string) string
}

type ConsoleFormatter struct{}

func (c ConsoleFormatter) direction(s string) string {
	return fmt.Sprintf("%s", brush.Red(s))
}

func (c ConsoleFormatter) place(s string) string {
	return fmt.Sprintf("%s", brush.Blue(s))
}

func (c ConsoleFormatter) time(s string) string {
	return fmt.Sprintf("%s", brush.Green(s))
}

type WebFormatter struct{}

func (w WebFormatter) direction(s string) string {
	return fmt.Sprintf("<span style=\"color:red;\">%s</span>", s)
}

func (w WebFormatter) place(s string) string {
	return fmt.Sprintf("<span style=\"color:blue;\">%s</span>", s)
}

func (w WebFormatter) time(s string) string {
	return fmt.Sprintf("<span style=\"color:green;\">%s</span>", s)
}
