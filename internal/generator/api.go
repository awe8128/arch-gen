package generator

func NewAPI() (string, string) {
	filename := "main.go"
	template := `
	package main
	func main() {
		ctx := context.Background()
		config, err := config.Load(".")
		if err != nil {
			log.Printf("Failed to load config: %v", err)
		}

		store, err := db.Init(ctx, config)
		if err != nil {
			log.Fatalf("Failed to load: %v", err)
		}

		handler := server.NewHandler(store)

		server, err := server.NewServer(config, handler)
		if err != nil {
			log.Fatalf("Failed to init server: %v", err)
		}

		server.Run(config.API_PORT)
	}
	`

	return template, filename
}
