package main

import "fmt"

func main() {
	var course1 [3]string
	//var course2 [4]string

	course1[0] = "go"
	course1[1] = "go1"
	course1[2] = "go2"

	for index, value := range course1 {
		fmt.Println(index, value)
	}

	//	init array
	//course2 := [3]string{"go", "grpc", "gin"}
	course2 := [3]string{2: "gin"}
	for index, value := range course2 {
		fmt.Println(index, value)
	}
	course3 := [...]string{"go"}

}
