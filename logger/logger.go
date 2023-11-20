package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var (
	logs zerolog.Logger
)

func GRPCLogger() zerolog.Logger {
	f, err := os.OpenFile("./logger/logs/grpc_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening grpc logs file: %v", err)
	}

	logs = zerolog.New(f).With().Timestamp().Logger()
	return logs
}

func HTTPLogger() zerolog.Logger {
	f, err := os.OpenFile("./logger/logs/http_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening http logs file: %v", err)
	}

	logs = zerolog.New(f).With().Timestamp().Logger()
	return logs
}

func ErrorWarningLogger() zerolog.Logger {
	f, err := os.OpenFile("./logger/logs/error_warning_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening http logs file: %v", err)
	}

	logs = zerolog.New(f).With().Timestamp().Logger()
	return logs
}

func DebugLogger() zerolog.Logger {
	f, err := os.OpenFile("./logger/logs/debug_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening http logs file: %v", err)
	}

	logs = zerolog.New(f).With().Timestamp().Logger()
	return logs
}
