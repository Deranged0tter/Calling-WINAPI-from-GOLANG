// example running MessageBoxW from user32.dll
package main

import (
	"fmt"
	"strconv"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	messageBoxText  string = "Test Text"
	messageBoxTitle string = "Test Title"
)

func main() {
	user32_DLL := windows.NewLazyDLL("user32.dll")

	MessageBoxW := user32_DLL.NewProc("MessageBoxW")
	r, _, err := MessageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(messageBoxText))),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(messageBoxTitle))),
		windows.MB_OKCANCEL,
	)

	fmt.Println("Message: ", strconv.FormatUint(uint64(*windows.StringToUTF16Ptr(messageBoxText)), 10))
	fmt.Println("Return Code: ", r)
	fmt.Println("Return Message: ", err)
}
