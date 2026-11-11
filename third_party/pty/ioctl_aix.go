//go:build aix
//+build aix

package pty

import (
	"syscall"
	"unsafe"
)

const (
	TIOCGWINSZ = syscall.TIOCGWINSZ
	TIOCSWINSZ = syscall.TIOCSWINSZ
)

//go:cgo_import_dynamic libc_ioctl ioctl "libc.a/shr_64.o"
//go:linkname procioctl libc_ioctl
var procioctl uintptr

// Defined in asm_aix_ppc64.s.
func sysvicall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

func ioctl(fd uintptr, cmd uintptr, args uintptr) error {
	if _, _, errno := sysvicall6(uintptr(unsafe.Pointer(&procioctl)), 3, fd, cmd, args, 0, 0, 0); errno != 0 {
		return errno
	}
	return nil
}
