package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info *log.Logger
	warning *log.Logger
	err *log.Logger
	writer io.Writer
}


func NewLogger(prefix string) *Logger {
	writter := io.Writer(os.Stdout)
	logger := log.New(writter, prefix, log.Ldate|log.Ltime)
	return &Logger{
		debug: log.New(writter, "DEBUG: ", logger.Flags()),
		info: log.New(writter, "INFO: ", logger.Flags()),
		warning: log.New(writter, "WARNING: ", logger.Flags()),
		err: log.New(writter, "ERROR : ", logger.Flags()),
		writer: writter,
	}

}


// NON FORMATED
func (l *Logger) Debug(v ...any){
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...any){
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...any){
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...any){
	l.err.Println(v...)
}

// ----- FORMATED ----------

func (l *Logger) Debugf(format string, v ...any){
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...any){
	l.info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...any){
	l.warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...any){
	l.err.Printf(format, v...)
}