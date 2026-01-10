package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/shared/templates"
)

func NewRouter() (string, string) {

	filename := "router.go"

	template := fmt.Sprintf(`
	%s

	func (s *Server) NewRoute() {
		// metrics.Init()
		router := gin.New()

		// logger := logx.New(s.config.LOGGER_LEVEL)

		router.Use(gin.Recovery())
		// router.Use(middleware.LoggingMiddleware(logger))
		// router.Use(middleware.CORSMiddleware())
		// router.Use(metrics.MetricsMiddleware())

		// _, err := schema.GetSwagger()
		// if err != nil {

		// 	return
		// }

		// schema.RegisterHandlers(router, &schema.ServerInterfaceWrapper{
		// 	Handler: s.handler,
		// })
		// Prometheus metrics endpoint
		// router.GET("/metrics", gin.WrapH(promhttp.Handler()))

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			})
		})
		s.engine = router
	}

	`, templates.NewPackageTemplate("server"))

	return template, filename
}
