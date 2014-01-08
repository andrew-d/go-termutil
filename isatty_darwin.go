// +build darwin freebsd

package termutil

import (
    "os"
    "syscall"
    "unsafe"
)

func Isatty(file *os.File) bool {
    var termios syscall.Termios

    _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, file.Fd(),
        uintptr(syscall.TIOCGETA),
        uintptr(unsafe.Pointer(&termios)),
        0,
        0,
        0)
    return err == 0
}
