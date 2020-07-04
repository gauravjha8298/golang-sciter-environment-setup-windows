package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	//create sciter window object
	win,_ :=  window.New(sciter.SW_MAIN|sciter.SW_TITLEBAR | sciter.SW_TOOL, nil)

	//set window title
	win.SetTitle("Simple Sciter World")

	
	win.Show()
	
	win.Run()
}

