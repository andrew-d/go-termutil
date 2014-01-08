// +build !windows,!linux,!darwin,!freebsd,!cgo

package isatty

import "os"

func Isatty(file *os.File) bool {
    panic("Not implemented")
}
