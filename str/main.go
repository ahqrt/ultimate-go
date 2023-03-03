package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//name := "golang"
	//bytes := []rune(name)
	//fmt.Println(len(bytes))
	//	格式化输出
	username := "ahqrt"
	age := 20
	out := "hello " + username
	fmt.Println(out)
	fmt.Printf("用户名: %s, 年龄: %d \n", username, age) // 常用，但是性能没有上面的好
	//	字符串拼接 通过string的builder 进行字符串拼接
	var builder strings.Builder
	builder.WriteString("用户名: ")
	builder.WriteString(username)
	builder.WriteString(" 年龄: ")
	builder.WriteString(strconv.Itoa(age))

	re := builder.String()
	fmt.Println(re)
	//	字符串常用方法

	fmt.Println(strings.Contains(re, "hello"))
	fmt.Println(strings.Count(re, "a"))
	fmt.Println(len(strings.Split(re, " ")))
}
