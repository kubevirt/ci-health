package output

import "io"

type Handler interface {
	SetPath(path string)
	io.Writer
}
