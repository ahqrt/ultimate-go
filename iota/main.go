package main

import "fmt"

/*
*
iota 特殊的常量，可以视为是一个编译器可以修改的常量
比如 正常定义常量
*/
//const (
//	Err1 = 1
//	Err2 = 2
//	Err3 = 4
//	Err4 = 4
//)

// 如果使用iota
const (
	Err1 = iota
	Err2
	Err3 = "haha"
	Err4
	Err5 = iota
)

/*
	如果内部中断了iota的使用，后续想要恢复必须显式声明
	内部自增默认是int类型
	iota简化了const类型的定义
*/

func main() {
	fmt.Println(Err1, Err2)
}
