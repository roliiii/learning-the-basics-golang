package main

import (
	"syscall"
	"unsafe"
)

func main() {
	kernel32DLL := syscall.NewLazyDLL("user32.dll")
	findWindowAProc := kernel32DLL.NewProc("FindWindowW")
	ret, _, _ := syscall.Syscall6(findWindowAProc.Addr(), 2, 0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Untitled - Notepad"))),
		0,
		0,
		0,
		0)

	procOpenProcess := kernel32DLL.NewProc("MessageBoxW")
	syscall.Syscall6(procOpenProcess.Addr(),
		4,
		ret,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Title"))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Message"))),
		0x00000004,
		0,
		0)
}
