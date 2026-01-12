package immutables

import (
	"fmt"
	"path/filepath"

	"github.com/awe8128/arch-gen/templates"
	"github.com/awe8128/arch-gen/utils/fs"
)

func GenerateServer(root string) error {
	path := filepath.Join(root, "presentation", "server")
	filename := "server.go"

	content := fmt.Sprintf(
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
	if err := fs.GenerateFile(content, path, filename); err != nil {
		return err
	}

	return nil
}
