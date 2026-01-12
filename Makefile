run:
	go run main.go -c arch.yaml
	
server:
	cd be && go run cmd/api/main.go

.PHONY:
	run server