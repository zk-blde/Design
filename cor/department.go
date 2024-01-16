package main

// 处理者接口
type Department interface {
	execute(*Patient)
	setNext(Department)
}
