package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultConfigFilePath = ".testconfig"
)

func (s *Service) getConfigFile() (*os.File, error) {
	if s.ConfigFilePath == "" {
		s.ConfigFilePath = defaultConfigFilePath
	}

	fp, err := filepath.Abs(s.ConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf("getting absolute path: %v", err)
	}

	fh, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0o755)
	if err != nil {
		return nil, fmt.Errorf("opening file: %v", err)
	}

	return fh, nil
}

func (s *Service) SetConfigFile(fpath string) {
	s.ConfigFilePath = fpath
}

func (s *Service) LoadConfig() {
	fh, err := s.getConfigFile()
	if err != nil {
		panic(err)
	}

	defer fh.Close()

	data, err := ioutil.ReadAll(fh)
	if err != nil {
		panic(err)
	}

	if len(data) < 2 {
		return
	}

	if err = json.Unmarshal(data, &s.store); err != nil {
		panic(err)
	}
}

func (s *Service) SaveConfig() {
	data, err := json.MarshalIndent(s.store, "", "  ")
	if err != nil {
		panic(err)
	}

	fh, err := s.getConfigFile()
	if err != nil {
		panic(err)
	}

	defer fh.Close()

	fh.Truncate(0)
	fh.Write(data)
}

func (s *Service) loopSaveConfig() {
	for {
		s.SaveConfig()
		time.Sleep(time.Second)
	}
}
