package abstract

// CanBackUp describes any service which can yield backup data as a byte slice
type CanBackUp interface {
	BackupKey() string
	Backup() ([]byte, error)
}

// CanRestore describes any service which can be restored from a byte slice
type CanRestore interface {
	BackupKey() string
	IsBackupKey(string) bool
	Restore([]byte) error
}

// BackupRestoreManager describes a service that would manage all services that
// can backup and restore
type BackupRestoreManager interface {
	BackupAll() map[string][]byte
	RestoreAll(map[string][]byte) map[string]error
}
