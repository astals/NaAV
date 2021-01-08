package main

import (
	"./modes"
)

func main() {
	// check admin privileges
	// TODO: save config file
	// TODO: create uninstall $het (control panel)
	// TODO: grab command install / uninstall / only checks
	modes.Install()
}
