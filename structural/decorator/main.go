package main

import (
	"fmt"
	"io"

	"github.com/huytran2000-hcmus/design-patterns-in-go/structural/decorator/datasource/decorator"

	"github.com/huytran2000-hcmus/design-patterns-in-go/structural/decorator/datasource"
)

func main() {
	var ds datasource.DataSource
	ds, err := datasource.NewFileDataSource()
	defer func() {
		err = ds.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}

	ds = decorator.NewEncryptionDataSource(ds)
	ds = decorator.NewCompressionSource(ds)

	data := []byte("In the beginning was the Word")
	writeData(ds, data)
	readData(ds)

	data = []byte(", and the Word was with God")
	writeData(ds, data)
	readData(ds)

	data = []byte(", and the Word was God.")
	writeData(ds, data)
	readData(ds)
	fmt.Println()
}

func readData(ds datasource.DataSource) {
	data := make([]byte, 256)
	for {
		n, err := ds.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", data[:n])
	}
}

func writeData(ds datasource.DataSource, data []byte) {
	_, err := ds.Write(data)
	if err != nil {
		panic(err)
	}
}
