package templates

import (
	"fmt"
)

func LogxTemplate() (string, string) {
	filename := "logx.go"
	template := fmt.Sprintf(`%s
	
func New(cflog string) *slog.Logger {
	level := setLogLevel(cflog)

	var levelVar slog.LevelVar

	levelVar.Set(level)

	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: &levelVar,
	})

	logger := slog.New(h)
	slog.SetDefault(logger)

	return logger
}

func setLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}

}
	`, Package("logx"))

	return filename, template
}

func HelperTemplate() (string, string) {
	filename := "helper.go"
	template := fmt.Sprintf(`%s
func GetLogger(c *gin.Context) *slog.Logger {
	val, ok := c.Get(LoggerKey)
	if !ok {
		return slog.Default()
	}

	logger, ok := val.(*slog.Logger)
	if !ok {
		return slog.Default()
	}

	return logger
}

func ContextWithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, l)
}

// Wrapper for each layers
func From(ctx context.Context) *slog.Logger {
	if v := ctx.Value(LoggerKey); v != nil {
		if l, ok := v.(*slog.Logger); ok {
			return l
		}
	}
	return slog.Default()
}

func Info(ctx context.Context, msg string, args ...any) {
	logger := From(ctx)
	logger.Info(msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	logger := From(ctx)
	logger.Error(msg, args...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	logger := From(ctx)
	logger.Debug(msg, args...)
}

	
	`, Package("logx"))

	return filename, template
}

func CustomCodeTemplate() (string, string) {
	filename := "consts.go"
	template := fmt.Sprintf(`%s

type customCode int

const LoggerKey string = "logger"

const (
	NotFound      customCode = 4004
	InternalError customCode = 5000
)
	
	`, Package("logx"))

	return filename, template
}
