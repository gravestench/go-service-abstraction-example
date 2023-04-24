package backup_restore

func (s *Service) BackupAll() map[string][]byte {
	backup := make(map[string][]byte)

	for key, b := range s.backupers {
		data, err := b.Backup()
		if err != nil {
			s.log.Error().Msgf("could not backup '%s': %v", key, err)
			continue
		}

		backup[key] = data
	}

	return backup
}
