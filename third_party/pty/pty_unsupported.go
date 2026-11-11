//go:build !linux && !darwin && !freebsd && !dragonfly && !netbsd && !openbsd && !solaris && !aix
// +build !linux,!darwin,!freebsd,!dragonfly,!netbsd,!openbsd,!solaris,!aix

package pty

import (
	"os"
)

func open() (pty, tty *os.File, err error) {
	return nil, nil, ErrUnsupported
}
