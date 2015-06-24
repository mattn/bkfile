// +build windows

package main

import (
	"syscall"
	"unsafe"
)

var (
	kernel32     = syscall.MustLoadDLL("kernel32.dll")
	procCopyFile = kernel32.MustFindProc("CopyFileW")
)

func copyFile(src, dst string) error {
	psrc, err := syscall.UTF16PtrFromString(src)
	if err != nil {
		return err
	}
	pdst, err := syscall.UTF16PtrFromString(dst)
	if err != nil {
		return err
	}
	r, _, err := procCopyFile.Call(uintptr(unsafe.Pointer(psrc)), uintptr(unsafe.Pointer(pdst)), uintptr(0))
	if r == 0 {
		return err
	}
	return nil
}
