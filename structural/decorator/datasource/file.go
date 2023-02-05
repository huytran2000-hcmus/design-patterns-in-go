package datasource

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const (
	fnamePattern = "data*.txt"
	dir          = ""
)

type FileDataSource struct {
	fname   string
	wf      *os.File
	rOffSet int64
}

func NewFileDataSource() (*FileDataSource, error) {
	file, err := os.CreateTemp(dir, fnamePattern)
	if err != nil {
		return nil, err
	}

	return &FileDataSource{fname: file.Name(), wf: file}, nil
}

func (fs *FileDataSource) Read(data []byte) (n int, err error) {
	line, err := fs.readLine()
	fs.rOffSet += int64(len(line))

	line = strings.TrimSuffix(line, "\n")
	copy(data, line)
	return len(line), err
}

func (fs *FileDataSource) Write(data []byte) (n int, err error) {
	n, err = fs.writeLine(data)
	if n > 0 {
		n = n - len("\n")
	}
	return n, err
}

func (fs *FileDataSource) Close() (err error) {
	defer func() {
		rErr := os.Remove(fs.fname)
		if rErr != nil {
			err = rErr
		}
	}()

	defer func() {
		wErr := fs.wf.Close()
		if wErr != nil {
			err = wErr
		}
	}()

	return nil
}

func (fs *FileDataSource) getReadFile() (*os.File, error) {
	file, err := os.Open(fs.fname)
	_, err = file.Seek(fs.rOffSet, io.SeekStart)
	return file, err
}

func (fs *FileDataSource) readLine() (string, error) {
	file, err := fs.getReadFile()
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			err = closeErr
		}
	}()
	if err != nil {
		return "", err
	}

	br := bufio.NewReader(file)
	line, err := br.ReadString('\n')
	return line, err
}

func (fs *FileDataSource) writeLine(data []byte) (int, error) {
	data = append(data, '\n')
	n, err := fs.wf.Write(data)
	return n, err
}
