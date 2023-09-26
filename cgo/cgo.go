package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: ./hello.a
#include <stdlib.h>
#include <hello.h>
*/
import "C"

import (
	"unsafe"
)

func main() {
	cStr := C.CString("hello world !") // golang 操作c 标准库中的CString函数
	C.print_fun1(cStr)                 // 调用C函数：print_fun 打印输出

	defer C.free(unsafe.Pointer(cStr)) // 因为 CSstring 这个函数没有对变量申请的空间进行内存释放
}
