//go:build windows
// +build windows

package utils

import "syscall"

func GetSysProcAttr() *syscall.SysProcAttr {
	return nil
}
