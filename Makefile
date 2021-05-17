 .PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server server/main.go
	bin/server

.PHONY: run-dev
run-dev:
	bash -c "export ENV=local && nodemon --exec go run src/main.go --signal SIGTERM"