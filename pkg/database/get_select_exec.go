package database

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

const (
	sqliteLocked           = "database is locked (5) (SQLITE_BUSY)"
	retryDelay             = 50 * time.Millisecond
	maximumRetryCount      = 10
	errLockedRetryExceeded = "sqlite database locked, maximum retry attempts exceeded"
)

func (s *Service) Get(result interface{}, query string, args ...interface{}) error {
	for attempt := 0; attempt < maximumRetryCount; attempt++ {
		err := s.db.Get(result, query, args...)
		if err == nil {
			return nil // happy path
		}

		if !strings.Contains(err.Error(), sqliteLocked) {
			return err // unhappy path
		}

		time.Sleep(retryDelay)
	}

	return errors.New(errLockedRetryExceeded) // unhappy path
}

func (s *Service) Select(result interface{}, query string, args ...interface{}) error {
	for attempt := 0; attempt < maximumRetryCount; attempt++ {
		err := s.db.Select(result, query, args...)
		if err == nil {
			return nil // happy path
		}

		if !strings.Contains(err.Error(), sqliteLocked) {
			return err // unhappy path
		}

		time.Sleep(retryDelay)
	}

	return errors.New(errLockedRetryExceeded) // unhappy path
}

func (s *Service) Exec(query string, args ...interface{}) (sql.Result, error) {
	for attempt := 0; attempt < maximumRetryCount; attempt++ {
		res, err := s.db.Exec(query, args...)
		if err == nil {
			return res, nil // happy path
		}

		if !strings.Contains(err.Error(), sqliteLocked) {
			return nil, err // unhappy path
		}

		time.Sleep(retryDelay)
	}

	return nil, errors.New(errLockedRetryExceeded) // unhappy path
}
