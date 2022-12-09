package model

type File struct {
	Name string
	Size int
}

func NewFile(name string, size int) *File {
	return &File{Name: name, Size: size}
}
