package datasource

type DataSource interface {
	Read(data []byte) (n int, err error)
	Write(data []byte) (n int, err error)
	Close() error
}
