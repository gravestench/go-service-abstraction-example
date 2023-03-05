package static_assets

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed embedded
var embedded embed.FS

func (s *Service) Open(name string) (fs.File, error) {
	f, err := embedded.Open(name)
	if err != nil {
		err = fmt.Errorf("opening file in embedded filesystem: %v", err)
	}

	return f, err
}

func (s *Service) ReadDir(name string) ([]fs.DirEntry, error) {
	list, err := embedded.ReadDir(name)
	if err != nil {
		err = fmt.Errorf("reading directory in embedded filesystem: %v", err)
	}

	return list, err
}

func (s *Service) ReadFile(name string) ([]byte, error) {
	data, err := embedded.ReadFile(name)
	if err != nil {
		err = fmt.Errorf("reading file from embedded filesystem: %v", err)
	}

	return data, err
}
