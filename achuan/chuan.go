package main

import (
	"context"
	"fmt"
)

type MyStruct struct {
	Value string
}

func modifyStruct(s *MyStruct) {

	*s = MyStruct{Value: "new value"}
}

func main() {
	// 创建一个 MyStruct 实例
	originalStruct := MyStruct{Value: "original value"}

	// 打印原始值
	fmt.Println("Before:", originalStruct.Value)

	// 将 MyStruct 的指针传递给 modifyStruct 函数
	modifyStruct(&originalStruct)

	// 打印修改后的值，应该看到 "new value"
	fmt.Println("After:", originalStruct.Value)
}

func testaaa(f func()) {
	if f != nil {
		fmt.Println("in")
		f()
	}
	fmt.Println("end")
}

func test123() func() {
	return func() {
		fmt.Println("test123")
	}
}

func IsContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
