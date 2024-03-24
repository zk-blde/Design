package main

import "fmt"

// 具体策略

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}
