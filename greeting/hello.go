package greeting

import (
	color "github.com/fatih/color"
)

var greeting string = color.BlueString("Golang for Brave!")

func Hello() string {
	return greeting
}
