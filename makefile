# note: call scripts from /scripts
gobuild:
	GOARCH=amd64 GOOS=linux go build -o build/main cmd/main/main.go
gorun: gobuild
	./build/main
gotest:
	go test -v ./...
gofix_modules:
	go mod vendor