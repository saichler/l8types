package ifs

type IStorage interface {
	Write(string, []byte) error
	Read(string) ([]byte, error)
}
