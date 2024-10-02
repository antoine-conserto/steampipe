//go:build darwin || linux
// +build darwin linux

package utils

import "syscall"

func GetSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setpgid:    true,
		Foreground: false,
	}
}
