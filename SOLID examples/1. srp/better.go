package srp

import (
	"io/fs"
	"os"
	"path"
)

type FileStorage struct {
	storagePath string
}

func (s *FileStorage) WriteFile(filename string, content []byte) {
	os.WriteFile(path.Join(s.storagePath, filename), []byte(content), fs.FileMode(0600))
}

type Entity struct {
	ID   string
	Name string
}

type Indexer struct {
	storage *FileStorage
}

func New(storage *FileStorage) *Indexer {
	return &Indexer{
		storage: storage,
	}
}

func (i *Indexer) Add(e Entity) {
	i.storage.WriteFile(e.Name, []byte(e.ID))
}
