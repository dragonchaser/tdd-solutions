package ocp

import (
	"io/fs"
	"os"
	"path"
)

type Storage interface {
	Persist(key, value string)
}

type FileStorage struct {
	storagePath string
}

func (s *FileStorage) Persist(key, value string) {
	os.WriteFile(path.Join(s.storagePath, key), []byte(value), fs.FileMode(0600))
}

type MemoryStorage struct {
	store map[string]string
}

func (s *MemoryStorage) Persist(key string, value string) {
	s.store[key] = value
}

type Entity struct {
	ID   string
	Name string
}

type Indexer struct {
	storage Storage
}

func New(storage Storage) *Indexer {
	return &Indexer{
		storage: storage,
	}
}

func (i *Indexer) Add(e Entity) {
	i.storage.Persist(e.Name, e.ID)
}
