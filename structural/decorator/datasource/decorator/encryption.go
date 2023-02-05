package decorator

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/huytran2000-hcmus/design-patterns-in-go/structural/decorator/datasource"
)

const secret = "6368616e676520746869732070617373"

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 0o5}

type EncryptionDataSource struct {
	ds datasource.DataSource
}

func NewEncryptionDataSource(ds datasource.DataSource) *EncryptionDataSource {
	return &EncryptionDataSource{ds: ds}
}

func (en *EncryptionDataSource) Read(data []byte) (n int, err error) {
	cipherText := make([]byte, len(data))
	n, err = en.ds.Read(cipherText)
	if err != nil {
		copy(data, cipherText)
		return n, err
	}
	cipherText = cipherText[:n]

	return readUnencrypted(data, cipherText)
}

func (en *EncryptionDataSource) Write(data []byte) (n int, err error) {
	cipherText := make([]byte, len(data))
	err = encryptData(cipherText, data)
	if err != nil {
		return 0, err
	}

	n, err = en.ds.Write(cipherText)
	return n, err
}

func encryptData(cipherText []byte, data []byte) error {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText, data)
	return nil
}

func (en *EncryptionDataSource) Close() error {
	return en.ds.Close()
}

func readUnencrypted(data []byte, cipherText []byte) (int, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return 0, err
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, cipherText)
	return len(cipherText), err
}
