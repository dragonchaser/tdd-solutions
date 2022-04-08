package srp

import (
	"io/fs"
	"os"
	"path"
)

type Entity struct {
	ID   string
	Name string
}

type Indexer struct {
	storagePath string
}

func New(path string) *Indexer {
	return &Indexer{
		storagePath: path,
	}
}

func (i *Indexer) Add(e Entity) {
	os.WriteFile(path.Join(i.storagePath, e.Name), []byte(e.ID), fs.FileMode(0600))
}
