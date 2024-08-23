package log

type Config struct {
	FilePath   string
	MaxSize    int
	MaxAge     int
	OutputType OutputType
}

type OutputType string

const (
	OutputTypeConsole OutputType = "console"
	OutputTypeFile    OutputType = "file"
)

const (
	MsgKey = "msg"
	ErrKey = "err"
)
