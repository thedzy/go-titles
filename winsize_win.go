//go:build windows
// +build windows

package main

import (
	"golang.org/x/sys/windows"
	"syscall"
)

// getWinSize Get the width of the window on Windows
func getWinSize() (int, error) {
	// Get window dimensions for Windows
	h, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return 0, err
	}
	var csbi windows.ConsoleScreenBufferInfo
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(h), &csbi); err != nil {
		return 0, err
	}

	return int(csbi.Size.X), nil
}
