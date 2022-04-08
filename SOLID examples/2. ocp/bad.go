package ocp

import (
	"io/fs"
	"os"
	"path"
)

type FileStorage struct {
	storagePath string
}

func (s *FileStorage) WriteFile(filename string, content []byte) {
	os.WriteFile(path.Join(s.storagePath, filename), content, fs.FileMode(0600))
}

type MemoryStorage struct {
	store map[string]string
}

func (s *MemoryStorage) Store(key string, value string) {
	s.store[key] = value
}

type Entity struct {
	ID   string
	Name string
}

type Indexer struct {
	storage interface{}
}

func New(storage interface{}) *Indexer {
	return &Indexer{
		storage: storage,
	}
}

func (i *Indexer) Add(e Entity) {
	switch s := i.storage.(type) {
	case *FileStorage:
		s.WriteFile(e.Name, []byte(e.ID))
	case *MemoryStorage:
		s.Store(e.Name, e.ID)
	}
}
