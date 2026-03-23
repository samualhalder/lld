package file

import (
	"github.com/samualhalder/lld/file_system/content"
	"github.com/samualhalder/lld/file_system/metadata"
)

type File struct {
	Name     string
	Ext      string
	Metadata *metadata.Metadata
	Content  *content.Content
}

func NewFile(name string) *File {
	metaData := &metadata.Metadata{}
	metaData.AssignCT()
	return &File{
		Name:     name,
		Metadata: metaData,
		Content:  &content.Content{},
	}
}
