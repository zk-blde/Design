package main

import "fmt"

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning 连接器插入 mac 机。")
}
