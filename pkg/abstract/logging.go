package abstract

import (
	"github.com/rs/zerolog"
)

// Logger is a generic logging service that steals its
// interface from zerolog because of the simplicity of using it.
type Logger interface {
	Trace() Messager
	Debug() Messager
	Info() Messager
	Warn() Messager
	Error() Messager
	Fatal() Messager
	SetLevel(int) Logger
	Logger() *zerolog.Logger
}

// Messager is something that can print or fprint a string.
// This is a pattern taken from zerolog, as well.
type Messager interface {
	Msg(string)
	Msgf(format string, v ...interface{})
}

type HasLogger interface {
	Logger() *zerolog.Logger
	SetLogger(Logger)
}
