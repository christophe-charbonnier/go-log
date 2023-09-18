package log

import (
	"testing"

	"github.com/rs/zerolog"
)

func Test(t *testing.T) {
	t.Run("colors", func(t *testing.T) {
		SetLevel(zerolog.InfoLevel)
		t.Run("no-color", func(t *testing.T) {
			SetColor(false)
			message := "without colors"
			Info().Msg(message)
			Log().Msg(message)
			Print(message)
		})
		t.Run("with-color", func(t *testing.T) {
			SetColor(true)
			message := "with colors"
			Info().Msg(message)
			Log().Msg(message)
			Print(message)
		})
	})

	t.Run("NoLevel", func(t *testing.T) {
		SetLevel(zerolog.NoLevel)
		Log().Msg("message with no level")
		Info().Msg("message with info level")
		Debug().Msg("message with debug level")
		Print("message without level and caller")
	})

	t.Run("InfoLevel", func(t *testing.T) {
		SetLevel(zerolog.InfoLevel)
		Log().Msg("message with no level")
		Info().Msg("message with info level")
		Debug().Msg("message with debug level")
		Print("message without level and caller")
	})

	t.Run("DebugLevel", func(t *testing.T) {
		SetLevel(zerolog.DebugLevel)
		Log().Msg("message with no level")
		Info().Msg("message with info level")
		Debug().Msg("message with debug level")
		Print("message without level and caller")
	})
}
