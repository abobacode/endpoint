package xerror

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrMaxRetriesReached  = errors.New("max retries reached")
	ErrEmptyPrograms      = errors.New("empty programs")
	ErrMaxIdleConns       = errors.New("max_idle_conns must be greater than zero")
	ErrMaxOpenConns       = errors.New("max_open_conns must be greater than zero")
	ErrMaxConnMaxLifetime = errors.New("max_conn_max_lifetime must be greater than zero")
	ErrTooLongTitle       = errors.New("title too long")
	ErrUnhandledTask      = errors.New("unhandled task type")
)

type WithoutNotificationError struct {
	err error
}

func (s *WithoutNotificationError) Error() string {
	return fmt.Sprintf("skip error: %s", s.err.Error())
}

func (s *WithoutNotificationError) Unwrap() error {
	return s.err
}

type SkipError struct {
	err error
}

func (s *SkipError) Error() string {
	return fmt.Sprintf("skip error: %s", s.err.Error())
}

func (s *SkipError) Unwrap() error {
	return s.err
}

type ReExecutableErr struct {
	err error
}

func (r *ReExecutableErr) Error() string {
	return fmt.Sprintf("re-executable error: %s", r.err.Error())
}

func (r *ReExecutableErr) Unwrap() error {
	return r.err
}
