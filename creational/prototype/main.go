package main

import (
	"fmt"
	"os"

	"github.com/huytran2000-hcmus/design-patterns-in-go/creational/prototype/file"
)

func main() {
	file1 := file.New("file1.txt")
	file2 := file.New("file2.txt")
	file3 := file.New("file3.txt")

	directory2 := file.NewDirectory("directory2", file2, file3)
	directory1 := file.NewDirectory("directory1", file1, directory2)
	directory1Clone := directory1.Clone()

	fmt.Println("File system tree: ")
	file.PrintFileTree(os.Stdout, "  ", directory1)
	fmt.Println("File system's Clone tree: ")
	file.PrintFileTree(os.Stdout, "  ", directory1Clone)
}
