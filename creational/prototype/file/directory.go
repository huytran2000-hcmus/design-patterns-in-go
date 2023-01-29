package file

type Directory struct {
	File
	children []Inode
}

func (d *Directory) GetTree(indentation string) string {
	var tree string
	tree += indentation + d.name
	tree += NewLine

	for _, child := range d.children {
		tree += child.GetTree(indentation + indentation)
		tree += NewLine
	}

	return tree
}

func (d *Directory) Clone() Inode {
	children := make([]Inode, len(d.children))

	for i, inode := range d.children {
		childClone := inode.Clone()
		children[i] = childClone
	}

	return NewDirectory(d.name+"_clone", children...)
}

func NewDirectory(name string, children ...Inode) *Directory {
	return &Directory{
		File:     File{name: name},
		children: children,
	}
}
