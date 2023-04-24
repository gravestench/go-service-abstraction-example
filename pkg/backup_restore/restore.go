package backup_restore

func (s *Service) RestoreAll(backups map[string][]byte) map[string]error {
	errors := make(map[string]error)

	for backupKey, backup := range backups {
		for _, r := range s.restorers {
			if !r.IsBackupKey(backupKey) {
				continue
			}

			if err := r.Restore(backup); err != nil {
				errors[backupKey] = err
			}
		}
	}

	if len(errors) == 0 {
		errors = nil
	}

	return errors
}
