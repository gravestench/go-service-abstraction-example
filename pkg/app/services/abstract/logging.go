package abstract

// Logger is a generic logging service that steals its
// interface from zerolog because of the simplicity of using it.
type Logger interface {
	Info() Messager
	Warn() Messager
	Error() Messager
	Fatal() Messager
}

// Messager is something that can print or fprint a string.
// This is a pattern taken from zerolog, as well.
type Messager interface {
	Msg(string)
	Msgf(format string, v ...interface{})
}
