package file

import (
	"io"
)

const NewLine = "\n"

type Inode interface {
	GetTree(indentation string) string
	Clone() Inode
}

func PrintFileTree(out io.Writer, indentation string, node Inode) {
	out.Write([]byte(node.GetTree(indentation)))
}
