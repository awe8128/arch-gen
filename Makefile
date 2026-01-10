run:
	go run main.go -c test.yaml
server:
	cd be && go run cmd/api/main.go

.PHONY:
	run server