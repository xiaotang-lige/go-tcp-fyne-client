package tool

import "os"

type path struct {
}

func (*path) Project() (path string) {
	path, _ = os.Getwd()
	return path
}
