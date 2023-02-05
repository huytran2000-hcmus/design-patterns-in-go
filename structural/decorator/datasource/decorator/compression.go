package decorator

import (
	"compress/flate"
	"io"

	"github.com/huytran2000-hcmus/design-patterns-in-go/structural/decorator/datasource"
)

const compressionLevel = -2

type CompressionDataSource struct {
	ds datasource.DataSource
	wc *flate.Writer
}

func NewCompressionSource(ds datasource.DataSource) *CompressionDataSource {
	return &CompressionDataSource{
		ds: ds,
	}
}

func (cp *CompressionDataSource) Read(data []byte) (n int, err error) {
	rc := flate.NewReader(cp.ds)
	defer func() {
		rErr := rc.Close()
		if rErr != nil && rErr != io.ErrUnexpectedEOF {
			err = rErr
		}
	}()
	n, err = rc.Read(data)

	if err == io.ErrUnexpectedEOF {
		return n, io.EOF
	}
	return n, err
}

func (cp *CompressionDataSource) Write(data []byte) (n int, err error) {
	err = cp.initWriterIfNotAlready()
	if err != nil {
		return 0, err
	}

	n, err = cp.wc.Write(data)
	defer func() {
		flushErr := cp.wc.Flush()
		if flushErr != nil {
			err = flushErr
		}
	}()

	return n, err
}

func (cp *CompressionDataSource) Close() (err error) {
	defer func() {
		dsErr := cp.ds.Close()
		if dsErr != nil {
			err = dsErr
		}
	}()

	defer func() {
		wErr := cp.wc.Close()
		if wErr != nil {
			err = wErr
		}
	}()

	return nil
}

func (cp *CompressionDataSource) initWriterIfNotAlready() error {
	if cp.wc == nil {
		var err error
		cp.wc, err = flate.NewWriter(cp.ds, compressionLevel)
		if err != nil {
			return err
		}
	}
	return nil
}
