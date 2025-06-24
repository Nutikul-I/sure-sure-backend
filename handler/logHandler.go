package handler

type LogHandler interface {
}

type logHandler struct {
}

func NewLogHandler() logHandler {
	return logHandler{}
}
