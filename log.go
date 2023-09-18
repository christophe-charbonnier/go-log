package log

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const (
	TimeFormat = "[2006/01/02 15:04:05]"
)

var (
	writer = &zerolog.ConsoleWriter{
		TimeFormat: TimeFormat,
		Out:        os.Stderr,
		NoColor:    false,
	}

	DefaultLogger = zerolog.New(writer).With().Timestamp().Caller().Logger()
)

func DisableColor() {
	writer.NoColor = true
	pterm.DisableColor()
}

func EnableColor() {
	writer.NoColor = false
	pterm.EnableColor()
}

func SetColor(b bool) {
	switch b {
	case true:
		EnableColor()
	case false:
		DisableColor()
	}
}

func init() {
	IsEnvDefined := func(s string) bool {
		_, found := os.LookupEnv(s)
		return found
	}

	SetColor(!IsEnvDefined("NO_COLOR"))

	switch {
	case IsEnvDefined("DEBUG"):
		SetLevel(zerolog.DebugLevel)
	default:
		SetLevel(zerolog.InfoLevel)
	}

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.CallerMarshalFunc = func(_ uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + fmt.Sprintf("%-3v", line)
	}
}

func Info() *zerolog.Event {
	return DefaultLogger.Info()
}

func Warn() *zerolog.Event {
	return DefaultLogger.Warn()
}

func Debug() *zerolog.Event {
	return DefaultLogger.Debug()
}

func Error() *zerolog.Event {
	return DefaultLogger.Error()
}

func Log() *zerolog.Event {
	w := *writer
	w.FormatLevel = func(i interface{}) string { return "" }
	// log := logger.Output(w)
	log := zerolog.New(w).With().Timestamp().Logger()
	return log.Log()
}

func Print(v ...any) {
	// w := *writer
	// w.FormatLevel = func(i interface{}) string { return "" }
	// log := zerolog.New(w).With().Timestamp().Logger()
	Log().Msg(fmt.Sprint(v...))
}

func Println(v ...any) {
	Log().Msg(fmt.Sprintln(v...))
}

func Printf(format string, v ...interface{}) {
	Print(fmt.Sprintf(format, v...))
}

func With() zerolog.Context {
	return DefaultLogger.With()
}

func GetLevel() zerolog.Level {
	return DefaultLogger.GetLevel()
}

func SetLevel(lvl zerolog.Level) {
	zerolog.SetGlobalLevel(lvl)
}
