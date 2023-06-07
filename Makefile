GOMOD=vendor

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/feed cmd/feed/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/feed-server cmd/feed-server/main.go	

lambda-server:
	if test -f main; then rm -f main; fi
	if test -f feed-server.zip; then rm -f feed-server.zip; fi
	GOOS=linux go build -mod $(GOMOD) -ldflags="-s -w" -o main cmd/feed-server/main.go
	zip feed-server.zip main
	rm -f main
