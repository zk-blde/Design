package main

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("适配器将 Lightning 信号转换为 USB。")
	w.windowMachine.insertIntoUSBPort()
}
