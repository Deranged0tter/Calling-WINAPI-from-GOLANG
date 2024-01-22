package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	path = "test.txt"
)

func main() {
	kernel32 := windows.NewLazyDLL("kernel32.dll")

	createFileW := kernel32.NewProc("CreateFileW")
	lpPath, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		log.Fatal(err)
	}
	r, _, err := createFileW.Call(
		uintptr(unsafe.Pointer(lpPath)),
		windows.GENERIC_READ,
		windows.FILE_SHARE_READ,
		0,
		windows.CREATE_NEW,
		0,
		0,
	)

	fmt.Println("Path: ", path)
	fmt.Println("Return Code: ", r)
	fmt.Println("Return Message: ", err)
}
