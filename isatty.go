// +build !windows,!linux,!darwin,!freebsd,!cgo

package termutil

import "os"

func Isatty(fd uintptr) bool {
    panic("Not implemented")
}
