package config_file_manager

import (
	"encoding/json"
)

func (s *Service) IsBackupKey(s2 string) bool {
	return s.BackupKey() == s2
}

func (s *Service) Restore(bytes []byte) error {
	s.mux.Lock()

	err := json.Unmarshal(bytes, &s.store)
	if err != nil {
		s.mux.Unlock()
		return err
	}

	s.mux.Unlock()

	s.SaveConfig()

	return nil
}

func (s *Service) BackupKey() string {
	return s.Name()
}

func (s *Service) Backup() ([]byte, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	return json.Marshal(s.store)
}
