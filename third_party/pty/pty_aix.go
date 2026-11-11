//go:build aix
//+build aix

package pty

import (
	"os"
	"strconv"
	"syscall"
)

func open() (pty, tty *os.File, err error) {
	p, err := os.OpenFile("/dev/ptc", os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		return nil, nil, err
	}
	// In case of error after this point, make sure we close the ptmx fd.
	defer func() {
		if err != nil {
			_ = p.Close() // Best effort.
		}
	}()

	sname, err := ptsname(p)
	if err != nil {
		return nil, nil, err
	}

	t, err := os.OpenFile(sname, os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		return nil, nil, err
	}
	return p, t, nil
}

func minor(x uint64) uint32 { return uint32(x) }

func ptsdev(fd uintptr) (uint64, error) {
        var status syscall.Stat_t
        if err := syscall.Fstat(int(fd), &status); err != nil {
                return 0, err
        }
        return uint64(minor(status.Rdev)), nil
}

func ptsname(f *os.File) (string, error) {
        dev, err := ptsdev(f.Fd())
        if err != nil {
                return "", err
        }
        fn := "/dev/pts/" + strconv.FormatInt(int64(dev), 10)

        if err := syscall.Access(fn, 0); err != nil {
                return "", err
        }
        return fn, nil
}
