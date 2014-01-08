// +build !windows,!linux,!darwin,!freebsd,!cgo

package termutil

import "os"

func Isatty(file *os.File) bool {
    panic("Not implemented")
}
