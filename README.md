# go-termutil

This package exposes some very baic, useful functions:

    Isatty(file *os.File) bool

It will return whether or not the given file is a TTY, attempting to use native operations
when possible.  It wil fall back to using the `isatty()` function from `unistd.h` through
cgo if on an unknown platform.
