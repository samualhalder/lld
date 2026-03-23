package directory

import (
	"fmt"

	"github.com/samualhalder/lld/file_system/file"
	"github.com/samualhalder/lld/file_system/metadata"
)

type Directory struct {
	Name        string
	Directories map[string]*Directory
	Files       map[string]*file.File
	ParentDir   *Directory
	Metadata    *metadata.Metadata
}

func NewDirectory(name string, dir *Directory) *Directory {
	metadata := &metadata.Metadata{}
	metadata.AssignCT()
	return &Directory{
		Name:        name,
		Directories: make(map[string]*Directory),
		Files:       make(map[string]*file.File),
		ParentDir:   dir,
		Metadata:    metadata,
	}
}

func (d *Directory) CreateDir(name string) error {
	_, exist := d.Directories[name]
	if exist {
		return fmt.Errorf("folder already exists")
	}
	newDir := NewDirectory(name, d)
	d.Directories[name] = newDir
	return nil
}

func (d *Directory) GoTo(name string) (*Directory, error) {
	chldDir, exists := d.Directories[name]
	if !exists {
		return nil, fmt.Errorf("no such directories")
	}
	return chldDir, nil
}

func (d *Directory) Back() *Directory {
	return d.ParentDir
}

func (d *Directory) CreateFile(name string) error {
	_, exists := d.Directories[name]
	if exists {
		return fmt.Errorf("file already exists")
	}
	newFile := file.NewFile(name)
	d.Files[name] = newFile
	return nil
}
