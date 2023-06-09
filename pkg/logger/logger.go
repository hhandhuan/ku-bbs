package logger

import (
	"fmt"
	"github.com/hhandhuan/ku-bbs/pkg/config"
	"github.com/hhandhuan/ku-bbs/pkg/utils"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var instance *zerolog.Logger

func GetInstance() *zerolog.Logger {
	return instance
}

func Initialize(config *config.Logger) {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return utils.ResolvePathFileName(file) + ":" + strconv.Itoa(line)
	}

	fmt.Println(123213)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano

	var output io.Writer

	output = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	fileLogger := &lumberjack.Logger{
		Filename:   config.Path,
		MaxSize:    config.MaxSize, // 5M
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)

	logger := zerolog.New(output).Level(zerolog.Level(config.Level)).With().Caller().Timestamp().Logger()

	instance = &logger
}
