package abstract

type Messager interface {
	Msg(string)
	Msgf(format string, v ...interface{})
}

type Logger interface {
	Info() Messager
	Warn() Messager
	Error() Messager
	Fatal() Messager
}
