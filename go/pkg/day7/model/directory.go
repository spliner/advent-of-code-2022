package model

type Directory struct {
	Name     string
	Parent   *Directory
	Children map[string]*Directory
	Files    []*File
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		Name:     name,
		Parent:   parent,
		Children: make(map[string]*Directory),
		Files:    make([]*File, 0),
	}
}

func (d *Directory) Size() int {
	var size int
	for _, file := range d.Files {
		size += file.Size
	}
	for _, directory := range d.Children {
		size += directory.Size()
	}
	return size
}

func (d *Directory) AddChild(child *Directory) {
	d.Children[child.Name] = child
}

func (d *Directory) FindChild(name string) (*Directory, bool) {
	child, ok := d.Children[name]
	return child, ok
}

func (d *Directory) AddFile(file *File) {
	d.Files = append(d.Files, file)
}
