package immutables

import (
	"fmt"

	"github.com/awe8128/arch-gen/templates"
)

func GenerateServer() (string, string) {
	filename := "server.go"

	template := fmt.Sprintf(
		`
		%s

	type Server struct {
		config  *config.Config
		handler *Handler
		engine  *gin.Engine
	}

	func NewServer(config *config.Config, handler *Handler) (*Server, error) {
		server := &Server{
			config:  config,
			handler: handler,
		}

		server.NewRoute()

		return server, nil
	}

	func (s *Server) Run(port string) error {
		err := s.engine.Run(port)
		if err != nil {
			return err
		}
		return nil
	}

		`,
		templates.Package("server"),
	)

	return template, filename
}
