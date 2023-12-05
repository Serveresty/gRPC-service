server-build:
	export GO111MODULE="on"; \
	go mod download; \
	go mod vendor; \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o cmd/server/main.go
deploy:
	rsync -avz cmd/server/main.go  hiro@server.com:/web-server/
restart-service:
	ssh hiro@server.com systemctl --user restart go-server-protei.service