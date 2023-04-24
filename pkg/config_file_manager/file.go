package config_file_manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

const (
	defaultConfigDir      = "~/.config/example"
	defaultConfigFileName = "config.json"
)

func (s *Service) getConfigFile() (*os.File, error) {
	if s.cfgFilePath == "" {
		s.cfgFilePath = defaultConfigFileName
	}

	fp, err := filepath.Abs(filepath.Join(s.DirPath(), s.cfgFilePath))
	if err != nil {
		return nil, fmt.Errorf("getting absolute path: %v", err)
	}

	if err = createDirsForFile(fp); err != nil {
		s.log.Fatal().Msgf("could not create config directory: %v", err)
	}

	if err = createFileIfNotExists(fp); err != nil {
		s.log.Fatal().Msgf("could not create config file: %v", err)
	}

	info, err := os.Stat(fp)
	if err != nil {
		return nil, fmt.Errorf("getting absolute path: %v", err)
	}

	s.lastModified = info.ModTime()

	fh, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0o755)
	if err != nil {
		return nil, fmt.Errorf("opening file: %v", err)
	}

	return fh, nil
}

func (s *Service) SetConfigFile(fpath string) {
	s.cfgFilePath = fpath
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
	s.mux.Lock()
	data, err := json.MarshalIndent(s.store, "", "  ")
	if err != nil {
		panic(err)
	}

	s.mux.Unlock()

	fh, err := s.getConfigFile()
	if err != nil {
		panic(err)
	}

	defer fh.Close()

	fh.Truncate(0)
	fh.Write(data)

	stat, err := fh.Stat()
	if err != nil {
		return
	}

	s.lastModified = stat.ModTime()
}

func (s *Service) loopSaveConfig() {
	for {
		if s.configHasDelta() {
			s.SaveConfig()
		}
		time.Sleep(time.Second)
	}
}

func (s *Service) configHasDelta() bool {
	if current, _ := json.Marshal(s.store); s.previous != string(current) {
		if s.previous == "" {
			s.previous = string(current)
			return false
		}

		s.previous = string(current)
		return true
	}

	return false
}

func createDirsForFile(filePath string) error {
	dirPath := filepath.Dir(filePath)
	return os.MkdirAll(dirPath, os.ModePerm)
}

func expandTilde(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return strings.Replace(path, "~", usr.HomeDir, 1), nil
}

func createFileIfNotExists(filePath string) error {
	_, err := os.Stat(filePath)

	// If the file does not exist, create it.
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	}

	return nil
}
