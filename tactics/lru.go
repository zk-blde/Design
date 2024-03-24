package main

import "fmt"

// 具体策略

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}
