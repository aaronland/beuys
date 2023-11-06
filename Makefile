GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/feed cmd/feed/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/feed-server cmd/feed-server/main.go	

lambda-server:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f feed-server.zip; then rm -f feed-server.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="-s -w" -tags lambda.norpc -o bootstrap cmd/feed-server/main.go
	zip feed-server.zip bootstrap
	rm -f bootstrap
