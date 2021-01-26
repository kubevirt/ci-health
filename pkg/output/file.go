package output

import "io/ioutil"

type FileHandler struct {
	path string
}

func NewFileHandler(path string) *FileHandler {
	return &FileHandler{
		path,
	}
}

func (f *FileHandler) SetPath(path string) {
	f.path = path
}

func (f *FileHandler) Write(p []byte) (int, error) {
	err := ioutil.WriteFile(f.path, p, 0644)

	if err != nil {
		return 0, err
	}

	return len(p), nil
}
