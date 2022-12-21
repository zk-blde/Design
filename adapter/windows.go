package main

import "fmt"

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB连接器插在windows机器上。")
}
