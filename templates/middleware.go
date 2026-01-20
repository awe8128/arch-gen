package templates

import "fmt"

func MiddlewareTemplate() (string, string) {

	filename := "middleware.go"
	template := fmt.Sprintf(`%s
	
func AuthMiddleware() {
	// Authentication middleware logic goes here
}

func LoggingMiddleware(base *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		requestLog := base.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
		)

		ctx := logx.ContextWithLogger(c.Request.Context(), requestLog)
		c.Request = c.Request.WithContext(ctx)

		c.Set(logx.LoggerKey, requestLog)
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		slog.Info("request completed",
			slog.Int("status", status),
			slog.Duration("latency", latency),
		)

	}
}

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3001", "http://127.0.0.1:3000", "http://127.0.0.1:3001"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	config.AllowCredentials = true
	return cors.New(config)
}
	
	
	`, Package("middleware"))

	return filename, template
}
