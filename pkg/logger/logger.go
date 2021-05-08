package logger

type Level int8

type Fields map[string]interface{}

const(
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic

)

func (l Level) String()  {

}
