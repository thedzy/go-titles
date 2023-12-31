//go:build linux || darwin
// +build linux darwin

package main

import (
	"golang.org/x/sys/unix"
	"os"
)

// getWinSize Get the full size of the window (*nix/mac)
func getWinSize() (int, error) {
	// Get window dimensions for Unix
	winDimensions, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0, err
	}
	return int(winDimensions.Col), nil
}
