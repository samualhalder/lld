package main

import (
	"fmt"
	"strings"

	"github.com/samualhalder/lld/file_system/directory"
)

type FileSystem struct {
	rootDir *directory.Directory
	currDir *directory.Directory
}

func (f *FileSystem) cd(str string) error {
	tmp := f.currDir
	if str == ".." {
		if f.currDir.ParentDir != nil {
			f.currDir = f.currDir.ParentDir
			return nil
		} else {
			return fmt.Errorf("no parent directory to go")
		}
	}
	cmds := strings.Split(str, "/")
	for _, cmd := range cmds {
		dir, have := f.currDir.Directories[cmd]
		if !have {
			f.currDir = tmp
			return fmt.Errorf("no such directories")
		}
		f.currDir = dir
	}
	return nil
}

func (f *FileSystem) mkdir(name string) error {
	return f.currDir.CreateDir(name)
}

func (f *FileSystem) touch(fileName string) error {
	return f.currDir.CreateFile(fileName)
}
