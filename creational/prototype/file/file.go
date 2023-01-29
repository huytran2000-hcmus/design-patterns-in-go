package file

type File struct {
	name string
}

func (f *File) GetTree(indentation string) string {
	return indentation + f.name + NewLine
}

func (f *File) Clone() Inode {
	return New(f.name + "_clone")
}

func New(name string) *File {
	return &File{name: name}
}
